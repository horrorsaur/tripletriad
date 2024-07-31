package main

func Sum(row any) int {
	var tmp int
	for _, v := range row.(Row) {
		tmp = tmp + v
	}
	return tmp
}
