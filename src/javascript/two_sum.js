/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
function twoSum(nums, target) {
  let store = {};

  for (let i = 0; i < nums.length; i++) {
    let item = nums[i];
    let diff = target - item;

    if (store[diff] === undefined) {
      store[item] = i;
    } else {
      return [store[diff], i];
    }
  }

  return [];
}

const items = [3, 2, 4];
const target = 6;

console.log(twoSum(items, target));
