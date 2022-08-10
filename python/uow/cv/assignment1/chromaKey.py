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
        print("Task one:", argument.option, argument.fileNames)

        # read original image
        original_image = cv2.imread(argument.fileNames[0])

        # color space conversion
        c1, c2, c3 = original_image, original_image, original_image
        match argument.option:
            case argument.optionXYZ:
                xyz_image = cv2.cvtColor(original_image, cv2.COLOR_BGR2XYZ)
                c1, c2, c3 = cv2.split(xyz_image)
            case argument.optionLab:
                lab_image = cv2.cvtColor(original_image, cv2.COLOR_BGR2LAB)
                c1, c2, c3 = cv2.split(lab_image)
            case argument.optionYCrCb:
                ycrcb_image = cv2.cvtColor(original_image, cv2.COLOR_BGR2YCrCb)
                c1, c2, c3 = cv2.split(ycrcb_image)
            case argument.optionHSB:
                hsb_image = cv2.cvtColor(original_image, cv2.COLOR_BGR2HSV)
                c2, c3, c1 = cv2.split(hsb_image)
            case _:
                pass

        # create a normal window
        cv2.namedWindow("Assignment1", cv2.WINDOW_NORMAL)

        # resize image down to fit maximum window size(1280*720)
        original_image = self.resize_image(original_image, int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)
        c1 = self.resize_image(c1, int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)
        c2 = self.resize_image(c2, int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)
        c3 = self.resize_image(c3, int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)

        # convert 2 dimension image to 3 dimension image which each channel has the same value
        c1 = np.stack((c1, c1, c1), axis=2)
        c2 = np.stack((c2, c2, c2), axis=2)
        c3 = np.stack((c3, c3, c3), axis=2)

        # build image window grid & show
        np_horizontal1 = np.hstack((original_image, c1))
        np_horizontal2 = np.hstack((c2, c3))
        np_vertical = np.vstack((np_horizontal1, np_horizontal2))

        # show & wait for destroy
        cv2.imshow("Assignment1", np_vertical)
        cv2.waitKey(0)
        cv2.destroyAllWindows()

    def get_task_type(self, argument):
        if argument.option in argument.validOptions \
                and argument.option != argument.optionUnknown\
                and len(argument.fileNames) == 1:
            return self.typeOne
        elif argument.option == argument.optionUnknown and len(argument.fileNames) == 2:
            return self.typeTwo
        return self.typeUnknown

    def process_task_two(self, argument):
        print("Task two:", argument.fileNames)

        # read scenic & green screen images
        scenic_image = cv2.imread(argument.fileNames[0])
        green_screen_image = cv2.imread(argument.fileNames[1])
        print(scenic_image.shape)
        print(green_screen_image.shape)

        # resize images to 640, 360
        scenic_image = self.resize_image(scenic_image, int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)
        green_screen_image = self.resize_image(green_screen_image, int(MAX_WIDTH/2), int(MAX_HEIGHT/2), cv2.INTER_AREA)
        print(scenic_image.shape)
        print(green_screen_image.shape)

        # create mask of green screen
        green_screen_image_hsv = cv2.cvtColor(green_screen_image, cv2.COLOR_BGR2HSV)
        upper = np.array([77, 255, 255])
        lower = np.array([35, 43, 46])
        mask = cv2.inRange(green_screen_image_hsv, lower, upper)

        # create white background image
        white_screen_image = green_screen_image.copy()
        fg_height = green_screen_image_hsv.shape[0]
        fg_width = green_screen_image_hsv.shape[1]
        for w in range(fg_width):
            for h in range(fg_height):
                if mask[h, w] == 255:
                    white_screen_image[h, w] = (255, 255, 255)

        # find contour
        mask_invert = cv2.bitwise_not(mask)
        contours, _ = cv2.findContours(mask_invert, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)

        # find bounding rect of the largest area
        largest_index = 0
        largest_area = 0
        for index in range(len(contours)):
            x, y, w, h = cv2.boundingRect(contours[index])
            if w * h > largest_area:
                largest_area = w * h
                largest_index = index

        # draw rect
        x, y, w, h = cv2.boundingRect(contours[largest_index])

        # clip mask image
        clip_image = white_screen_image[y:y+h, x:x+w]
        mask_clip_image = mask[y:y+h, x:x+w]

        # calc image dimension to align the image to the middle of the bg
        bg_height = scenic_image.shape[0]
        bg_width = scenic_image.shape[1]
        clip_image_in_scenic_x = int(bg_width / 2 - w / 2)
        clip_image_in_scenic_y = int(bg_height - h)
        print("bg height:", bg_height, "bg width:", bg_width)
        print("roi height:", h, "roi width:", w)
        print("scenic x:", clip_image_in_scenic_x, "scenic y:", clip_image_in_scenic_y)

        # create merged image
        merged_image = scenic_image.copy()
        top_left_x = clip_image_in_scenic_x
        top_left_y = clip_image_in_scenic_y
        for height in range(h):
            for width in range(w):
                if mask_clip_image[height, width] != 255:
                    merged_image[top_left_y+height, top_left_x+width] = clip_image[height, width]

        # build image window grid & show
        np_horizontal1 = np.hstack((green_screen_image, white_screen_image))
        np_horizontal2 = np.hstack((scenic_image, merged_image))
        np_vertical = np.vstack((np_horizontal1, np_horizontal2))

        # show & wait for destroy
        cv2.imshow("Assianment1", np_vertical)
        cv2.waitKey(0)
        cv2.destroyAllWindows()

    def resize_image(self, image, width, height, interpolation):
        # check params
        if width is None and height is None:
            return image

        # calculate ratio to scale image
        image_height = image.shape[0]
        image_width = image.shape[1]

        ratio = width / float(image_width)
        print("width ratio:", ratio)
        return cv2.resize(image, (width, int(image_height * ratio)), 0, 0, interpolation=interpolation)
        # if float(image_width) / image_height > float(1280)/720:
        #     ratio = width / float(image_width)
        #     print("width ratio:", ratio)
        #     return cv2.resize(image, (width, int(image_height * ratio)), 0, 0, interpolation=interpolation)
        # else:
        #     ratio = height / float(image_height)
        #     print("height ratio:", ratio)
        #     return cv2.resize(image, (int(image_width * ratio), height), 0, 0, interpolation=interpolation)


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

# def show_image(image):
#     """
#     Function to display image
#     :param image:
#     :return:
#     """
#     image = cv2.imread(filename)
#     show_image(image)
#     cv2.namedWindow('image', cv2.WINDOW_NORMAL)
#     cv2.imshow('image', image)
#     cv2.waitKey(0)
#     cv2.destroyAllWindows()
#
#
# def parse_and_run():
#     file_name = sys.argv[1]
#     main(file_name)


