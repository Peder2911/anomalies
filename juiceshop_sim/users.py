
import random
import os
from seleniumwire import webdriver


CHROMEDRIVER_BINARY = os.path.expanduser(os.getenv("CHROMEDRIVER_BINARY","~/.local/bin/chromedriver"))
JUICESHOP_URL = os.getenv("JUICESHOP_URL", "http://localhost:8000")

class User():
    def __init__(self):
        self._ip = ".".join(str(random.randint(1,256)) for _ in range(4))
        self.driver = webdriver.Chrome(CHROMEDRIVER_BINARY)
        self.driver.request_interceptor = self._set_fake_headers 

    def _set_fake_headers(self, request):
        request.headers["X-My-Fake-Ip"] = self._ip
        request.headers["X-Fake-User-Agent"] = "Simulator"

    def simulate(self):
        self.driver.get(JUICESHOP_URL)

if __name__ == "__main__":
    user = User()
    user.simulate()
