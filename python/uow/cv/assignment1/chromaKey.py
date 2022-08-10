import cv2
import sys
import matplotlib.pyplot as plt
import numpy as np

MAX_HEIGHT = 720
MAX_WIDTH = 1280


# task type definition
class Task:
    typeUnknown = 0
    typeOne = 1
    typeTwo = 2

    def get_task_type(self, argument):
        if argument.option in argument.validOptions \
                and argument.option != argument.optionUnknown\
                and len(argument.fileNames) == 1:
            return self.typeOne
        elif argument.option == argument.optionUnknown and len(argument.fileNames) == 2:
            return self.typeTwo
        return self.typeUnknown

    def process_task(self, argument):
        match self.get_task_type(argument):
            case Task.typeOne:
                self.process_task_one(argument)
            case Task.typeTwo:
                self.process_task_two(argument)
            case _:
                print("Task type unknown!")
                return

    def process_task_one(self, argument):
        # read original image
        original_image = cv2.imread(argument.fileNames[0])

        # color space conversion
        c1, c2, c3 = cv2.split(self.convert_color_space(original_image, argument))

        # create a normal window
        cv2.namedWindow("Assignment1", cv2.WINDOW_NORMAL)

        # resize image down to fit maximum window size(1280*720) 16:9
        original_image, c1, c2, c3 = self.resize_images(
            [original_image, c1, c2, c3], int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)

        # convert 2 dimension image to 3 dimension image which each channel has the same value
        c1 = np.stack((c1, c1, c1), axis=2)
        c2 = np.stack((c2, c2, c2), axis=2)
        c3 = np.stack((c3, c3, c3), axis=2)

        # show window & destroy
        self.show_image_grid("Assignment1", original_image, c1, c2, c3)

    def process_task_two(self, argument):
        # read scenic & green screen images
        scenic_image, green_screen_image = cv2.imread(argument.fileNames[0]), cv2.imread(argument.fileNames[1])

        # resize image down to fit maximum window size(1280*720) 16:9
        scenic_image, green_screen_image = self.resize_images(
            [scenic_image, green_screen_image], int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)

        # get mask of green screen image
        mask = self.get_image_hsv_mask(green_screen_image, np.array([77, 255, 255]), np.array([35, 43, 46]))

        # create white background image
        white_screen_image = self.fill_image_color_by_mask(green_screen_image, (255, 255, 255), mask)

        # find contours of mask image
        contours, _ = cv2.findContours(cv2.bitwise_not(mask), cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)

        # find bounding rect of the largest contour area
        clip_x, clip_y, clip_w, clip_h = self.find_bouding_rect_by_area(contours)

        # clip image & mask image using the largest contour bounding rect
        clip_image = white_screen_image[clip_y:clip_y+clip_h, clip_x:clip_x+clip_w]
        clip_image_mask = mask[clip_y:clip_y+clip_h, clip_x:clip_x+clip_w]

        # calc image dimension to align the image to the middle of the background
        scenic_h, scenic_w = scenic_image.shape[0], scenic_image.shape[1]
        clip_image_in_scenic_x = int(scenic_w / 2 - clip_w / 2)
        clip_image_in_scenic_y = int(scenic_h - clip_h)

        # create merged image
        merged_image = self.merge_image_by_mask(
            scenic_image, clip_image, clip_image_in_scenic_x, clip_image_in_scenic_y, clip_h, clip_w, clip_image_mask)

        # show window & destroy
        self.show_image_grid("Assignment1", green_screen_image, white_screen_image, scenic_image, merged_image)

    def convert_color_space(self, image, argument):
        match argument.option:
            case argument.optionXYZ:
                return cv2.cvtColor(image, cv2.COLOR_BGR2XYZ)
            case argument.optionLab:
                return cv2.cvtColor(image, cv2.COLOR_BGR2LAB)
            case argument.optionYCrCb:
                return cv2.cvtColor(image, cv2.COLOR_BGR2YCrCb)
            case argument.optionHSB:
                return cv2.cvtColor(image, cv2.COLOR_BGR2HSV)
            case _:
                return image

    def resize_images(self, images, width, height, interpolation):
        resized_images = []
        for _, image in enumerate(images):
            resized_images.append(self.resize_image(image, width, height, interpolation))
        return resized_images

    def resize_image(self, image, width, height, interpolation):
        # check params
        if width is None and height is None:
            return image

        # calculate ratio to scale image
        image_height = image.shape[0]
        image_width = image.shape[1]

        ratio = width / float(image_width)
        return cv2.resize(image, (width, int(image_height * ratio)), 0, 0, interpolation=interpolation)

    def show_image_grid(self, name, image1, image2, image3, image4):
        # build image window grid & show
        np_horizontal1 = np.hstack((image1, image2))
        np_horizontal2 = np.hstack((image3, image4))
        np_vertical = np.vstack((np_horizontal1, np_horizontal2))

        # show & wait for destroy
        cv2.imshow(name, np_vertical)
        cv2.waitKey(0)
        cv2.destroyAllWindows()

    def get_image_hsv_mask(self, image, upper, lower):
        mask_image_hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)
        return cv2.inRange(mask_image_hsv, lower, upper)

    def fill_image_color_by_mask(self, src, color, mask):
        dst = src.copy()
        height = src.shape[0]
        width = src.shape[1]
        for h in range(height):
            for w in range(width):
                if mask[h, w] == 255:
                    dst[h, w] = color
        return dst

    def merge_image_by_mask(self, bg, fg, x, y, h, w, mask):
        bg_image = bg.copy()
        for height in range(h):
            for width in range(w):
                if mask[height, width] != 255:
                    bg_image[y + height, x + width] = fg[height, width]
        return bg_image

    def find_bouding_rect_by_area(self, contours):
        largest_index = 0
        largest_area = 0
        for index in range(len(contours)):
            x, y, w, h = cv2.boundingRect(contours[index])
            if w * h > largest_area:
                largest_area = w * h
                largest_index = index
        return cv2.boundingRect(contours[largest_index])


# argument type definition
class Argument:
    optionUnknown = ""
    optionXYZ = "-XYZ"
    optionLab = "-Lab"
    optionYCrCb = "-YCrCb"
    optionHSB = "-HSB"
    validOptions = (optionXYZ, optionLab, optionYCrCb, optionHSB)

    def __init__(self):
        self.option = self.optionUnknown
        self.fileNames = []

    def parse(self, args):
        """
        :param args: system arguments without first python file name
        :return: task type, option & file names
        """
        if len(args) > 2:
            print("Command line argument invalid! length > 2")
            return Task.typeUnknown, "", ""

        self.option = Argument.optionUnknown
        for arg in args:
            if arg in Argument.validOptions:
                self.option = arg
            else:
                self.fileNames.append(arg)


def main(args):
    # parse arguments
    argument = Argument()
    argument.parse(args)

    # process task
    task = Task()
    task.process_task(argument)


if __name__ == '__main__':
    main(sys.argv[1:])
