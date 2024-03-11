import logging
import os
import time
from dataclasses import dataclass, field

from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

from chrome_driver import get_chrome_driver
from processor import Processor

# Selectors
USERNAME_SELECTOR = "session_key"
PASSWORD_SELECTOR = "session_password"
SIGNIN_BUTTON_SELECTOR = "sign-in-form__submit-btn--full-width"
LINKEDIN_BASE_URL = "https://www.linkedin.com"
LINKEDIN_USERNAME = os.environ["LINKEDIN_USERNAME"]
LINKEDIN_PASSWORD = os.environ["LINKEDIN_PASSWORD"]
LINKEDIN_SAVED_POSTS_URL = "https://www.linkedin.com/my-items/saved-posts/"
SAVED_POSTS_SELECTOR = "workflow-results-container"
SAVED_POSTS_LIST_SELECTOR = "reusable-search__result-container"
SAVED_POST_SELECTOR = "entity-result__content-inner-container"
SAVED_POST_LINK_SELECTOR = "app-aware-link"


logger = logging.getLogger(__name__)


@dataclass
class LinkedinScraper:
    chrome_driver: webdriver.Chrome
    scraped_posts: list = field(default_factory=list)

    def login(self):
        self.chrome_driver.get(LINKEDIN_BASE_URL)
        # Wait up to 10 seconds for the elements to become available
        print("Waiting for USERNAME_SELECTOR...")
        WebDriverWait(self.chrome_driver, 10).until(
            EC.presence_of_element_located((By.ID, USERNAME_SELECTOR))
        )
        time.sleep(3)

        # Locate the username and password fields
        print("Locating USERNAME_SELECTOR and PASSWORD_SELECTOR...")
        username_input = self.chrome_driver.find_element(By.ID, USERNAME_SELECTOR)
        password_input = self.chrome_driver.find_element(By.ID, PASSWORD_SELECTOR)

        # Write the username and password
        username_input.send_keys(LINKEDIN_USERNAME)
        password_input.send_keys(LINKEDIN_PASSWORD)

        # Locate the sign-in button and click it
        print("Locating SIGNIN_BUTTON_SELECTOR...")
        sign_in_button = self.chrome_driver.find_element(
            By.CLASS_NAME, SIGNIN_BUTTON_SELECTOR
        )
        sign_in_button.click()

        time.sleep(20)
        print("Logged in!!!")

    def navigate_to_saved_posts(self):
        self.chrome_driver.get(LINKEDIN_SAVED_POSTS_URL)
        WebDriverWait(self.chrome_driver, 10).until(
            EC.presence_of_element_located((By.CLASS_NAME, SAVED_POSTS_SELECTOR))
        )

    def get_saved_posts(self) -> list | None:
        self.navigate_to_saved_posts()
        self.scroll_to_bottom()
        html_source = self.chrome_driver.page_source
        parser = BeautifulSoup(html_source, "html.parser")
        saved_posts = parser.find_all(class_=SAVED_POSTS_LIST_SELECTOR)
        if not saved_posts:
            logger.warning("No saved posts found")
            return
        for saved_post in saved_posts:
            try:
                saved_post_info = saved_post.find(class_=SAVED_POST_SELECTOR)
                saved_post_link = saved_post_info.find(class_=SAVED_POST_LINK_SELECTOR)
                if not saved_post_link:
                    saved_post_link = saved_post_info.find("a")
                if not saved_post_link:
                    logger.warning("No saved post link found")
                    continue

                saved_post_link = saved_post_link.__dict__["attrs"]["href"]
                saved_post_link = saved_post_link.split("?")[0]
                self.scraped_posts.append(saved_post_link)

            except Exception as e:
                logger.error("Error getting saved post info: %s", e)
                raise

        return self.scraped_posts

    def scroll_to_bottom(self):
        scroll_script = """
            function scrollDown() {
              window.scrollTo(0, document.body.scrollHeight);
            }

            // Store it in a variable in order to clear the interval later
            let intervalId = window.setInterval(scrollDown, 3000);
            
            return intervalId;
            """

        interval_id = self.chrome_driver.execute_script(scroll_script)
        time.sleep(300)

        clear_interval_script = f"window.clearInterval({interval_id})"
        self.chrome_driver.execute_script(clear_interval_script)


if __name__ == "__main__":
    chrome_driver = get_chrome_driver()
    linkedin_scraper = LinkedinScraper(chrome_driver=chrome_driver)
    linkedin_scraper.login()
    saved_posts = linkedin_scraper.get_saved_posts()
    if saved_posts:
        processor = Processor(unprocessed_links=saved_posts)
        processor.process()
    logger.info("Scraped %s posts", len(saved_posts))
