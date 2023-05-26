from random import choice
import argparse
import time
from os import getcwd, listdir, path
from subprocess import run
import darkdetect

parser = argparse.ArgumentParser(
                    prog='RandBackground',
                    description='Chooses a (pseudo-)random background between intervals')

parser.add_argument('-p', '--path', default=getcwd(), type=str)
parser.add_argument('-i', '--interval', default=300, type=int)

args = parser.parse_args()


class Backgrounds:
    def __init__(self, background_dir: str, interval_seconds: int) -> None:
        self.background_dir = background_dir
        self.interval_seconds = interval_seconds
        self.random_background_loop()

    def set_background(self, absolute_background_path: str):
        if darkdetect.isDark():
            uri_mode = "picture-uri-dark"
        else:
            uri_mode = "picture-uri"

        cmd = f"gsettings set org.gnome.desktop.background {uri_mode} {absolute_background_path}"
        run(cmd, shell=True, check=True)


    def choose_random_background(self):
        pictures = [file for file in listdir(self.background_dir)]
        random_background = choice(pictures)
        relative_background_path = "".join((self.background_dir, random_background))
        absolute_background_path = path.abspath(relative_background_path)
        self.set_background(absolute_background_path)

    def random_background_loop(self) -> None:
        while True:
            self.choose_random_background()
            time.sleep(self.interval_seconds)


if __name__ == "__main__":
    main_obj = Backgrounds(
        background_dir = args.path,
        interval_seconds = args.interval)
