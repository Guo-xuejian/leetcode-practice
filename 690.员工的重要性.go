/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	// map 方便查询
	employeeMap := make(map[int]*Employee)
	// 遍历写入对应指针节省空间
	for _, employee := range employees {
		employeeMap[(*employee).Id] = employee
	}
	return dfs(id, &employeeMap)
}

func dfs(id int, employeeMap *map[int]*Employee) (total int) {
	employee := (*employeeMap)[id]
	// 当前员工重要性
	total = (*employee).Importance
	// 直系员工重要性
	for _, employeeId := range (*employee).Subordinates {
		total += dfs(employeeId, employeeMap)
	}
	return
}

// @lc code=end

