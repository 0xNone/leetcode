def isBadVersion(version) -> bool:
    if version >= 4:
        return True
    
# class Solution:
#     def firstBadVersion(self, n):
#         """
#         :type n: int
#         :rtype: int
#         """
#         i_cache = n / 2
#         n = int(i_cache)
#         trueCloseNum = 0
#         falseCloseNum = 0
#         while True:
#             if isBadVersion(n):
#                 trueCloseNum = n
#                 n -= int(i_cache)
#             else:
#                 falseCloseNum = n
#                 n += int(i_cache)
#             if trueCloseNum - 1 == falseCloseNum:
#                 return trueCloseNum
#             if i_cache < 1:
#                 i_cache = 1
#             else:
#                 i_cache /= 2
    
# 网友思路
class Solution:
    def firstBadVersion(self, end):
        """
        :type n: int
        :rtype: int
        """
        start = 0
        while start < end:
            mid = start + (end - start) // 2
            if isBadVersion(mid):
                end = mid
            else:
                start = mid + 1
        return start

if __name__ == '__main__':
    s = Solution()
    print(6//2)
    print(s.firstBadVersion(5))