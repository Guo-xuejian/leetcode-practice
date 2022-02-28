#
# @lc app=leetcode.cn id=1603 lang=python3
#
# [1603] 设计停车系统
#

# @lc code=start
class ParkingSystem:

    def __init__(self, big: int, medium: int, small: int):
        self.big = big
        self.medium = medium
        self.small = small
        self.curr_big = 0
        self.curr_medium = 0
        self.curr_small = 0


    def addCar(self, carType: int) -> bool:
        if carType == 1 and self.curr_big < self.big:
            self.curr_big += 1
        elif carType == 2 and self.curr_medium < self.medium:
            self.curr_medium += 1
        elif carType == 3 and self.curr_small < self.small:
            self.curr_small += 1
        else:
            # 没有空余车位
            return False
        return True



# Your ParkingSystem object will be instantiated and called as such:
# obj = ParkingSystem(big, medium, small)
# param_1 = obj.addCar(carType)
# @lc code=end

