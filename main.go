package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var todos []map[string]string

func main() {
	for {
		clearScreen()
		fmt.Println("Daftar Pilihan:")
		fmt.Println("1. Menambahkan todo")
		fmt.Println("2. Menampilkan todo")
		fmt.Println("3. Selesaikan todo")
		fmt.Println("4. Menghapus todo")
		fmt.Println("5. Keluar")

		fmt.Print("Masukan Pilihan : ")
		input := inputUser()
		clearScreen()

		switch input {
		case "1":
			fmt.Print("Masukan todo : ")
			input = inputUser()
			addTodo(input)
		case "2":
			listTodos(todos)
		case "3":
			listTodos(todos)
			fmt.Print("\nSelesaikan todo no : ")
			input = inputUser()
			finishedTodo(parseIndex(input))
		case "4":
			listTodos(todos)
			fmt.Print("\nHapus todo no : ")
			input = inputUser()
			deleteTodo(parseIndex(input))
		case "5":
			fmt.Println("keluar...")
			return
		default:
			fmt.Println("Pilihan yang anda masukan salah")
		}

		fmt.Print("\nTekan enter untuk kembali ke menu utama")
		inputUser()

	}
}

// fungsi untuk mengambil input user
func inputUser() string {
	reader := bufio.NewReader(os.Stdin) //membuat reader yang akan membaca input dari os.Stdin
	input, _ := reader.ReadString('\n') //membaca input sampai baris baru
	input = strings.TrimSpace(input)    //menghilangkan new line atau spasi yang ada di akhir input

	return input
}

// fungsi untuk membersihkan layar
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// fungsi untuk menambahkan todo
func addTodo(todo string) {
	todos = append(todos, map[string]string{"todo": todo, "status": "belum selesai"})

	fmt.Println("Todo berhasil ditambahkan")
}

// fungsi untuk menampilkan list todo
func listTodos(todo []map[string]string) {
	if len(todo) == 0 {
		fmt.Println("Tidak ada daftar todo")
		return
	}

	fmt.Println("Daftar Todo:")
	fmt.Printf("%-5s %-30s %-5s\n", "No", "Todo", "Status")
	for i, todo := range todo {
		fmt.Printf("%-5d %-30s %-5s\n", i+1, todo["todo"], todo["status"])
	}
}

// fungsi untuk menyelesaikan todo
func finishedTodo(no int) {
	if no >= 0 && no < len(todos) {
		todos[no]["status"] = "selesai"
		fmt.Printf("Todo nomor %d berhasil diselesaikan\n", no+1)
	} else {
		fmt.Println("Nomor to-do tidak sesuai")
	}

}

// fungsi untuk menghapus todo
func deleteTodo(no int) {
	if no >= 0 && no < len(todos) {
		todos = append(todos[:no], todos[no+1:]...)
		fmt.Printf("Todo nomor %d berhasil dihapus\n", no+1)
	} else {
		fmt.Println("Nomor to-do tidak sesuai")
	}
}

// fungsi untuk memparsing index dari string ke int
func parseIndex(index string) int {
	var idx int
	fmt.Sscan(index, &idx)

	return idx - 1
}
