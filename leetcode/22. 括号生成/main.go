package main

var results []string

func generateParenthesis(n int) []string {

}

func backtrace(n int, open, close int, path string) {
	if len(path) == 2*n {
		p := path
		results = append(results, p)
	}

	if path > n {

	}
}
