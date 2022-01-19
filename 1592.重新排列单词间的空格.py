#
# @lc app=leetcode.cn id=1592 lang=python3
#
# [1592] 重新排列单词间的空格
#

# @lc code=start
# class Solution:
#     def reorderSpaces(self, text: str) -> str:
#         space_num = 0
#         word_list = []
#         word = ""
#         for one in text:
#             if one == " ":
#                 space_num += 1
#                 if word != "":
#                     word_list.append(word)
#                     word = ""
#             else:
#                 word += one
#         # 末尾单词无空格处理
#         if word != "":
#             word_list.append(word)
#         words_num = len(word_list)      # 单词数量
#         interval_num = words_num - 1    # 间隔数
#         space_interval_num = int(space_num / interval_num) if interval_num != 0 else 0  # 间隔处空格数
#         space_left_num = space_num % interval_num if interval_num != 0 else space_num           # 结尾处空格数
#         text = ""
#         for i in range(words_num):
#             text += word_list[i]    # 写入单词
#             if i == words_num-1:    # 结尾处写入剩余空格
#                 text += " "*space_left_num
#                 break
#             # 间隔处写入空格
#             text += " "*space_interval_num
#         return text

class Solution:
    def reorderSpaces(self, text: str) -> str:
        space_num = 0
        word_list = []
        word = ""
        for one in text:
            if one != " ":
                word += one
                print(word)
            else:
                space_num += 1
                if len(word) > 0:
                    word_list.append(word)
                    word = ""
        if len(word) > 0:
            word_list.append(word)
        space_interval_num = len(word_list) - 1
        space_interval = " "
        space_left = " "
        space_interval = space_interval * \
            int(space_num / space_interval_num) if space_interval_num != 0 else ""
        space_left = space_left * \
            (space_num % space_interval_num) if space_interval_num != 0 else space_num * space_left
        text = ""
        for i in range(space_interval_num + 1):
            text += word_list[i]
            if i == space_interval_num:
                text += space_left
                break
            text += space_interval
        return text
# @lc code=end
