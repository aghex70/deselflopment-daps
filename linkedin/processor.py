import csv
import logging
from dataclasses import dataclass, field
from datetime import datetime

from utils import generate_uuid

logger = logging.getLogger(__name__)


@dataclass
class Processor:
    unprocessed_links: list = field(default_factory=list)
    file_name: str = f'saved_posts_{datetime.now().strftime("%Y%m%d%H%M%S")}'

    def generate_file(self) -> None:
        with open(f"posts/{self.file_name}.csv", "w", newline="") as file:
            writer = csv.writer(file)
            for link in self.unprocessed_links:
                processed_link = self.process_link(link)
                writer.writerow([processed_link["name"], processed_link["link"]])

    def process_link(self, link: str) -> dict:
        return {
            "name": generate_uuid(),
            "link": link,
        }

    def process(self) -> None:
        self.generate_file()
        logger.info("File generated successfully")
