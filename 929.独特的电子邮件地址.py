#
# @lc app=leetcode.cn id=929 lang=python3
#
# [929] 独特的电子邮件地址
#

# @lc code=start
class Solution:
    def numUniqueEmails(self, emails: List[str]) -> int:
        emailSet = set()
        for email in emails:
            index = email.index("@")
            local_name = email[:index].split("+", 1)[0]
            local_name = local_name.replace(".", "")
            emailSet.add(local_name + email[index:])
        return len(emailSet)
# @lc code=end

