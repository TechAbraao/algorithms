def twoSum(nums, target):  
    for index, value in enumerate(nums):
        idxs = 1
        if (value + nums[idxs]) == target:
            return [index, idxs]
        idxs = idxs + 1


if __name__ == "__main__":
    # r1 = twoSum(nums=[2,7,11,15], target=9)
    # r2 = twoSum(nums=[3, 2, 4], target=6)
    # r3 = twoSum(nums=[3, 3], target=6)
    r4 = twoSum(nums=[3, 2, 3], target=6)
    # print(r1)
    # print(r2)
    # print(r3)
    print(r4)