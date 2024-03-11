from selenium import webdriver
from selenium.webdriver.chrome.options import Options as ChromeOptions
from selenium.webdriver.chrome.service import Service as ChromeService


def get_chrome_driver() -> webdriver.Chrome:
    chrome_service = ChromeService()
    driver = webdriver.Chrome(options=get_chrome_options(), service=chrome_service)
    driver.maximize_window()
    return driver


def get_chrome_options() -> ChromeOptions:
    chrome_options = ChromeOptions()
    chrome_options.page_load_strategy = "none"
    # chrome_options.add_argument("--headless=new")
    chrome_options.add_argument("--enable-automation")
    chrome_options.add_argument("--start-maximized")
    # chrome_options.add_argument(f"--window-size={width},{height}")
    chrome_options.add_argument("--lang=en-GB")
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-setuid-sandbox")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_options.add_argument("--disable-gpu")
    chrome_options.add_argument("--disable-accelerated-2d-canvas")
    # chrome_options.add_argument("--proxy-server='direct://")
    # chrome_options.add_argument("--proxy-bypass-list=*")
    chrome_options.add_argument("--allow-running-insecure-content")
    chrome_options.add_argument("--disable-web-security")
    chrome_options.add_argument("--disable-client-side-phishing-detection")
    chrome_options.add_argument("--disable-notifications")
    chrome_options.add_argument("--mute-audio")
    chrome_options.add_argument("--ignore-certificate-errors")
    chrome_options.add_argument("--remote-allow-origins=*")

    # Disable downloads
    chrome_options.add_experimental_option(
        "prefs",
        {
            "safebrowsing.enabled": "false",
            "download.prompt_for_download": False,
            "download.default_directory": "/dev/null",
            "download_restrictions": 3,
            "profile.default_content_setting_values.notifications": 2,
        },
    )
    return chrome_options
