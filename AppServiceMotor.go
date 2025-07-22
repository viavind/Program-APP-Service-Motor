package main

import "fmt"

const N int = 1000

type SparePart struct {
	Name     string
	SerialNo string
	Price    int
	Quantity int
}

type Customer struct {
	Name  string
	Phone string
	Plate string
}

type Transaction struct {
	CustomerName Customer
	Date         string
	SpareParts   SparePart
	TotalAmount  float64
}

type ServiceMotorApp struct {
	SpareParts    [N]SparePart
	Customers     [N]Customer
	Transactions  [N]Transaction
	TransactionID int
}

func display() {
	fmt.Println("--------------------------------")
	fmt.Println("=    Aplikasi Service Motor    =")
	fmt.Println("=      Pelayanan Service       =")
	fmt.Println("=       by Kelompok 13         =")
	fmt.Println("--------------------------------")
}

func menu() {
	fmt.Println("-------------------------------------------")
	fmt.Println("=   Menu : 1. Data SparePart         	   =")
	fmt.Println("=          2. Data Pelanggan              =")
	fmt.Println("=          3. Data Transaksi              =")
	fmt.Println("=          4. Hitung Tarif Service        =")
	fmt.Println("=          5. Cari Data Pelanggan         =")
	fmt.Println("=          6. Tampilkan Daftar SparePart  =")
	fmt.Println("=          7. Urutkan Daftar SparePart    =")
	fmt.Println("=          9. Untuk Keluar                =")
	fmt.Println("-------------------------------------------")
}

func menuSparePart() {
	fmt.Println("-------------------------------------------")
	fmt.Println("=   Menu : 1. Tambah Data SparePart       =")
	fmt.Println("=          2. Edit Data SparePart         =")
	fmt.Println("=          3. Hapus Data SparePart        =")
	fmt.Println("=          9. Keluar                      =")
	fmt.Println("-------------------------------------------")
}

func menuPelanggan() {
	fmt.Println("-------------------------------------------")
	fmt.Println("=   Menu : 1. Tambah Data Pelanggan       =")
	fmt.Println("=          2. Edit Data Pelanggan         =")
	fmt.Println("=          3. Hapus Data Pelanggan        =")
	fmt.Println("=          9. Keluar				       =")
	fmt.Println("-------------------------------------------")
	
}
func menuTranksaksi() {
	fmt.Println("-------------------------------------------")
	fmt.Println("=   Menu : 1. Tambah Data Tranksaksi      =")
	fmt.Println("=          2. Edit Data Tranksaksi        =")
	fmt.Println("=          3. Hapus Data Tranksaksi       =")
	fmt.Println("=          9. Keluar					   =")
	fmt.Println("-------------------------------------------")
}
func menuQuantity(){
	fmt.Println("-------------------------------------------")
	fmt.Println("=   Menu : 1. Paling Jarang Diganti       =")
	fmt.Println("=          2. Paling Sering Diganti       =")
	fmt.Println("=          9. Keluar                      =")
	fmt.Println("-------------------------------------------")
}

func menucaripelanggan(){
	fmt.Println("-------------------------------------------")
	fmt.Println("=   Menu : 1. Berdasarkan Tanggal         =")
	fmt.Println("=          2. Berdasarkan Spare Part      =")
	fmt.Println("=          9. Keluar                      =")
	fmt.Println("-------------------------------------------")
}

func mainmenu() {
	var pilih, i int
	var app ServiceMotorApp
	var tgl string
	var Customers Customer
	var serialNo, Name, nama, sparepart string
	menu()
	fmt.Print("Silakan pilih : ");fmt.Scan(&pilih)
	for pilih != 9{
		
	if pilih == 1 {
		for pilih != 9{
			menuSparePart()
			fmt.Print("Apa yang akan dilakukan? ");fmt.Scan(&pilih)
			if pilih == 1{
				addSparePart(&app, &i)
			}else if pilih == 2{
				editSparePart(&app, N, serialNo)
				showSparePart(&app,&i)
			}else if pilih == 3{
				deleteSparePart(&app, &i, serialNo)
				showSparePart(&app, &i)
			}
		}

	}else if pilih == 2{
		for pilih != 9{
			menuPelanggan()
			fmt.Print("Apa yang akan dilakukan? ");fmt.Scan(&pilih)
			if pilih == 1{
				addCustomer(&app, &i)
				}else if pilih == 2{
					editCustomers(&app, N, Name)
					showCustomer(&app, &i)
			}else if pilih == 3{
				deleteCostumer(&app, &i, Name)
				showCustomer(&app, &i)
			}
		}
		}else if pilih == 3{
			for pilih != 9{
				menuTranksaksi()
				fmt.Print("Apa yang akan dilakukan? ");fmt.Scan(&pilih)
				if pilih == 1{
					addTransaction(&app, &i)
					showTransaction(&app, &i)
				}else if pilih == 2{
					editTransactions(&app, i, Customers)
					showTransaction(&app, &i)
				}else if pilih == 3{
					deleteTransaction(&app, &i, Customers)
					showTransaction(&app, &i)
				}
			}
		}else if pilih == 4{
			fmt.Println(servicePrice(&app, N, nama))
		}else if pilih == 5{
		for pilih != 9 {
			menucaripelanggan()
			fmt.Print("Apa yang akan dilakukan? ");fmt.Scan(&pilih)
			if pilih == 1 {
				fmt.Println(searchPelangganbyDate(&app, N, tgl))
			}else if pilih == 2{
				fmt.Println(searchPelangganbySparePart(&app, N, sparepart))
			}
		}
		}else if pilih == 6{
			showSparePart(&app, &i)
		}else if pilih == 7{
			for pilih != 9{
				menuQuantity()
				fmt.Print("Apa yang akan dilakukan? ");fmt.Scan(&pilih)
				if pilih == 1{
					sortByQuantity(&app, i)
				}else if pilih == 2{
					sortByQuantityselection(&app,i)
				}
			}
		}
		menu()
		fmt.Print("Silakan pilih : ");fmt.Scan(&pilih)
		}
	}


		
func addSparePart(app *ServiceMotorApp, N *int) {
	fmt.Println("input data yang akan ditambah: ")
	fmt.Println("(No Seri Spare Part) (Nama Spare Part) (Quantity) (Harga)")
	fmt.Scan(&app.SpareParts[*N].SerialNo)
	for app.SpareParts[*N].SerialNo != "XXXX" {
		fmt.Scan(&app.SpareParts[*N].Name, &app.SpareParts[*N].Quantity, &app.SpareParts[*N].Price)
		*N++
		fmt.Scan(&app.SpareParts[*N].SerialNo)
	}
	fmt.Println("Maaf data yang Anda inputkan Tidak Valid")
}

func showSparePart(app *ServiceMotorApp, i *int){
	for j := 0; j < *i ; j++{
		fmt.Println(app.SpareParts[j].SerialNo, app.SpareParts[j].Name, app.SpareParts[j].Quantity, app.SpareParts[j].Price)
	}
}

func addCustomer(app *ServiceMotorApp, i *int) {
	fmt.Println("input data yang akan ditambah: ")
	fmt.Println("(Nama Pelanggan) (No Telepon) (Plat Nomor Kendaraan)")
	fmt.Scan(&app.Customers[*i].Name)
	for app.Customers[*i].Name != "XXXX" {
		fmt.Scan(&app.Customers[*i].Phone, &app.Customers[*i].Plate)
		*i++
		fmt.Scan(&app.Customers[*i].Name)
	}
	fmt.Println("Maaf data yang Anda inputkan Tidak Valid")
}

func showCustomer(app *ServiceMotorApp, i *int){
	for j := 0; j < *i; j++{
		fmt.Println(app.Customers[j].Name, app.Customers[j].Phone, app.Customers[j].Plate)
	}
}

func addTransaction(app *ServiceMotorApp, i *int) {
	fmt.Println("input data yang akan ditambah: ")
	fmt.Println("(Nama Pelanggan)  (Tanggal Transaksi)  (Nama Spare Part)  (Total Transaksi)") 
	fmt.Scan(&app.Transactions[*i].CustomerName.Name)
	for app.Transactions[*i].CustomerName.Name != "XXXX" {
		fmt.Scan(&app.Transactions[*i].Date, &app.Transactions[*i].SpareParts.Name, &app.Transactions[*i].TotalAmount)
		*i++
		fmt.Scan(&app.Transactions[*i].CustomerName.Name)
	}
	fmt.Println("Maaf data yang Anda inputkan Tidak Valid")
}

func showTransaction(app *ServiceMotorApp, i *int){
	for j := 0; j < *i; j++{
		fmt.Println(app.Transactions[j].CustomerName.Name, app.Transactions[j].Date, app.Transactions[j].SpareParts.Name, app.Transactions[j].TotalAmount)
	}
}

func deleteSparePart(app *ServiceMotorApp, N *int, serialNo string) {
	var temp, i int
	fmt.Print("Masukkan Nomor Serial yang akan dihapus: ");fmt.Scan(&serialNo)
	for i = 0; i < *N; i++ {
		if app.SpareParts[i].SerialNo == serialNo {
			temp = i
			app.SpareParts[i].Quantity--
		}
	}

	for temp < *N-1 {
		app.SpareParts[temp] = app.SpareParts[temp+1]
		temp++
	}
	*N--
}

func deleteCostumer(app *ServiceMotorApp, N *int, Name string) {
	var temp, i int
	fmt.Print("Masukkan nama yang akan dihapus: ");fmt.Scan(&Name)
	for i = 0; i < *N; i++ {
		if app.Customers[i].Name == Name {
			temp = i
		}
	}

	for temp < *N-1 {
		app.Customers[temp] = app.Customers[temp+1]
		temp++
	}
	*N--
}

func deleteTransaction(app *ServiceMotorApp, N *int, Customers Customer){
	var temp, i int
	fmt.Print("Masukan nama yang akan dihapus: ");fmt.Scan(&Customers.Name)
	for i = 0; i < *N; i++ {
		if  app.Transactions[i].CustomerName == Customers{
			temp = i
			app.Transactions[i].SpareParts.Quantity--
		}
	}

	for temp < *N-1 {
		app.Transactions[temp].CustomerName = app.Transactions[temp+1].CustomerName
	}
	*N--
}

func servicePrice(app *ServiceMotorApp, N int, nama string) float64 {
	var price float64
	fmt.Println("Total siapa ?")
	fmt.Scan(&nama)
	for i := 0; i < N; i++ {
		if app.Transactions[i].CustomerName.Name == nama {
			price += app.Transactions[i].TotalAmount
		}
	}

	return price
}

func searchPelangganbyDate(app *ServiceMotorApp, N int, tgl string) string {
	var found bool = false
	var i int = 0
	fmt.Print("Tanggal berapa yang anda cari? ")
	fmt.Scan(&tgl)
	for i < N {
		if app.Transactions[i].Date == tgl {
			found = true
			break
		}
		i++
	}
	if found {
		return app.Transactions[i].CustomerName.Name
	}
	return "Tidak ada pelanggan pada tanggal tersebut"
}

func searchPelangganbySparePart(app *ServiceMotorApp, N int, sparepart string) string {
	var left, right, mid int
	fmt.Print("Sparepart apa yang anda cari? ")
	fmt.Scan(&sparepart)
	left = 0
	right = N - 1

	for left <= right && app.Transactions[mid].SpareParts.Name != sparepart {
		mid = (left + right) / 2

		if sparepart < app.Transactions[mid].SpareParts.Name {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if app.Transactions[mid].SpareParts.Name == sparepart {
		return app.Transactions[mid].CustomerName.Name
	} else {
		return "Tidak ada pelanggan yang membeli sparepart tersebut"
	}
}

func sortByQuantity(app *ServiceMotorApp, N int) {
	for i := 1; i < N; i++ {
		key := app.SpareParts[i]
		j := i - 1
	
		for j >= 0 && app.SpareParts[j].Quantity < key.Quantity {
			app.SpareParts[j+1] = app.SpareParts[j]
			j--
		}
	
		app.SpareParts[j+1] = key
	}
	
	fmt.Println("Daftar SparePart Terurut berdasarkan Jumlah yang Paling jarang Diganti:")
	for i := 0; i < N; i++ {
		fmt.Println("Nama:", app.SpareParts[i].Name)
		fmt.Println("Serial Number:", app.SpareParts[i].SerialNo)
		fmt.Println("Harga:", app.SpareParts[i].Price)
		fmt.Println("Jumlah:", app.SpareParts[i].Quantity)
		fmt.Println("-----------------------------")
	}
}

func sortByQuantityselection(app *ServiceMotorApp, N int) {
	for i := 0; i < N-1; i++ {
		minIndex := i
	
		for j := i + 1; j < N; j++ {
			if app.SpareParts[j].Quantity < app.SpareParts[minIndex].Quantity {
				minIndex = j
			}
		}
	
		if minIndex != i {
			temp := app.SpareParts[i]
			app.SpareParts[i] = app.SpareParts[minIndex]
			app.SpareParts[minIndex] = temp
		}
	}	
	
	fmt.Println("Daftar SparePart Terurut berdasarkan Jumlah yang Sering Diganti:")
	for i := 0; i < N; i++ {
		fmt.Println("Nama:", app.SpareParts[i].Name)
		fmt.Println("Serial Number:", app.SpareParts[i].SerialNo)
		fmt.Println("Harga:", app.SpareParts[i].Price)
		fmt.Println("Jumlah:", app.SpareParts[i].Quantity)
		fmt.Println("-----------------------------")
	}
}


func editSparePart(app *ServiceMotorApp, N int, serialNo string){
	var newName, newSerialNo string
	var newPrice, newQuantity int
	fmt.Print("masukan nomor serial spare part yang akan diubah : "); fmt.Scan(&serialNo)
    for i:= 0;i<N;i++  {
        if app.SpareParts[i].SerialNo == serialNo  {
            fmt.Println("Masukkan data setelah diedit:")
			fmt.Println("(Nomor Serial) (Spare Part) (Stok) (Harga)")
			fmt.Scan( &newSerialNo, &newName, &newQuantity ,&newPrice)
			app.SpareParts[i].Name = newName
            app.SpareParts[i].SerialNo = newSerialNo
			app.SpareParts[i].Quantity = newQuantity
			app.SpareParts[i].Price = newPrice
        }
    }
}

func editCustomers(app *ServiceMotorApp, N int, name string){
	var newName, newPhone, newPlate string
	fmt.Print("masukan nama pelanggan yang akan diubah : "); fmt.Scan(&name)
    for i := 0; i < N; i++{
        if app.Customers[i].Name == name {
            fmt.Println("Masukkan data setelah diedit:")
			fmt.Println("(Nama Pelanggan) (Nomor Telefon) (Plat Nomor)")
			fmt.Scan(&newName, &newPhone, &newPlate)
			app.Customers[i].Name = newName
            app.Customers[i].Phone = newPhone
			app.Customers[i].Plate = newPlate
        }
    }
}

func editTransactions(app *ServiceMotorApp, N int, Customers Customer){
	var newDate string
	var newCustomerName Customer
	var newSpareParts SparePart
	var newTotalAmount float64
	fmt.Print("masukan pelanggan transaksi yang akan diubah : "); fmt.Scan(&Customers.Name)
    for i := 0; i < N; i++{
        if app.Transactions[i].CustomerName == Customers{
            fmt.Println("Masukkan buku setelah diedit:")
			fmt.Println("(Nama Pelanggan) (Tanggal Transaksi) (Nama Spare Part) (Total Pembayaran)")
			fmt.Scan(&newCustomerName.Name, &newDate, &newSpareParts, &newTotalAmount)
			app.Transactions[i].CustomerName.Name = newCustomerName.Name
            app.Transactions[i].Date = newDate
			app.Transactions[i].SpareParts.Name = newSpareParts.Name
			app.Transactions[i].TotalAmount = newTotalAmount
        }
    }
}

func main(){
	display()
	mainmenu()
}