package go_say_hello

/*
	Major Update
	- Biasanya kita merubah kode yg tidak backward compatibility, artinya org yg memanggil module kita akan terkena error yg signifikan.
	- Salah satu caranya dengan merubah nama module, biasanya ditambah versi di belakang nama module sebelum.
*/

// Add parameter string to SayHello function (major update)
func SayHello(name string) string {
	return "Hello World!"
}
