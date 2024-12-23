package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Product представляет продукт
type Product struct {
	ID          int
	ImageURL    string
	Name        string
	Description string
	Price       float64
	Articul		string
}

// Пример списка продуктов
var products = []Product{
	{ID: 1, ImageURL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.L-ZQ0V8yKQqOhi6RDA_tTQHaE7%26pid%3DApi&f=1&ipt=319c3a7b61f91869c8cc1b397e349b0e7880fc10cd325c1e5306dba721603044&ipo=images", Name: "Тюльпаны", Description: "Тюльпа́н — род многолетних травянистых луковичных растений семейства Лилейные, в современных систематиках включающий более 80 видов.", Price: 120, Articul: "2893702"},
	{ID: 2, ImageURL: "https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fdachnaya-zhizn.ru%2Fimages%2Fdacha%2Fgolovna-1.jpg&f=1&nofb=1&ipt=fa90346231b9774390d031819da8daba76a73bdae92892d0ce2d4c40b0fce955&ipo=images", Name: "Розы", Description: "Ро́за — собирательное название видов и сортов представителей рода Шипо́вник, выращиваемых человеком и растущих в дикой природе. Бо́льшая часть сортов роз получена в результате длительной селекции путём многократных повторных скрещиваний и отбора. Некоторые сорта являются формами дикорастущих видов.", Price: 130, Articul:"1787351"},
	{ID: 3, ImageURL: "https://agro-market24.ru/upload/medialibrary/303/btwgg3vuz7uk9a7vs3zvgqjvkl3o55qi.jpg", Name: "Лилии", Description: "Ли́лия — род растений семейства Лилейные. Многолетние травы, снабжённые луковицами, состоящими из мясистых низовых листьев, расположенных черепитчато, белого, розоватого или желтоватого цвета.", Price: 140, Articul:"87164187"},
	{ID: 4, ImageURL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Frastenievod.com%2Fwp-content%2Fuploads%2F2017%2F04%2F3-17.jpg&f=1&nofb=1&ipt=2066e10ba3ebfc9d70c4cc9e44d0e7f2f0eb3ba33198815c4ddcaf5078fa328f&ipo=images", Name: "Ромашки", Description: "Рома́шка — род однолетних цветковых растений семейства астровые, или сложноцветные, по современной классификации объединяет около 70 видовневысоких пахучих трав, цветущих с первого года жизни.", Price: 160, Articul:"86231896"},
	{ID: 5, ImageURL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fgreensotka.ru%2Fwp-content%2Fuploads%2F2019%2F12%2F2-raznovidnost-fialki-senpolija.jpg&f=1&nofb=1&ipt=5bf80c8cafa44f55ec563816d8bd72bb1dfd027134008d91e35a79750cc63fc1&ipo=images", Name: "Фиалки", Description: "Однолетник, у которого заложение зачатков стеблей происходит в год, предшествующий их росту. После созревания тонкие плодоножки опускают корзинки-соцветия к земле. Таким образом, зрелые семена лежат под кустом.", Price: 100, Articul:"9813649"},
	{ID: 6, ImageURL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fflo.discus-club.ru%2Fimages%2Fstories%2Ffoto-cvetov%2Fpeony%2Fpeony1.jpg&f=1&nofb=1&ipt=4ca2ab6b51c47e78af9ec6f615ef1b2b01e3390b0b5be58b25331df395c04da6&ipo=images", Name: "Пионы", Description: "Пио́н — род травянистых многолетников и листопадных кустарников. Единственный род семейства Пионовые, ранее род относили к семейству Лютиковых. Пионы цветут в конце весны, ценятся садоводами за пышную листву, эффектные цветы и декоративные плоды.", Price: 380, Articul:"43223687"},
	{ID: 7, ImageURL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimg.7dach.ru%2Fimage%2F1200%2F00%2F00%2F48%2F2017%2F02%2F28%2F49cb68.jpg&f=1&nofb=1&ipt=bea0afa32038da81e043d3a24fd0cb961fe9d9d377b9cf945b5d78571ae5fb2f&ipo=images", Name: "Маки", Description: "Маки — потрясающе красивые растения, покорившие сердца многих садоводов. Но не все так просто. Среди этих пламенных красавцев есть такие виды, которые запрещено выращивать в частных садах. Давайте разбираться вместе.", Price: 200, Articul:"41337060"},
	{ID: 8, ImageURL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fdari-cvety.com%2Fassets%2Fimages%2Fproducts%2F1347%2F5-sinih-gipsofil-2.jpeg&f=1&nofb=1&ipt=084209b26043d02a2e6d2ebd9e00c67fff1ad1699f8fa35ff2411ef0efa31a3f&ipo=images", Name: "Гипсофилы", Description: "Такое травянистое растение как гипсофила (Gypsophila) еще именуют качим, перекати-поле, гипсолюбка.", Price: 230, Articul:"81273447"},
	{ID: 9, ImageURL: "https://duckduckgo.com/i/b15a6fbd4875f984.jpg", Name: "Люпины", Description: "Люпи́н, или волчий боб — род растений из семейства Бобовые. Представлен однолетними и многолетними травянистыми растениями, полукустарничками, полукустарниками, кустарниками.", Price: 160, Articul:"89370190"},
	{ID: 10, ImageURL: "https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fmegaogorod.com%2Ffiles%2Ffield%2Fimage%2Fu2169%2F2_251.jpg&f=1&nofb=1&ipt=653d4d89807910266d58a34b0631556c95c21704aa87995842dbbbd1bd5835b9&ipo=images", Name: "Орхидеи", Description: "Орхи́дные, или Ятры́шниковые, также Орхиде́и — крупнейшее семейство однодольных растений. Более чем для 10 % представителей семейства характерен CAM-фотосинтез. Орхидные — древнее семейство, появившееся в позднемеловую эпоху. ", Price: 145, Articul:"59247291"},
	
	
}

// обработчик для GET-запроса, возвращает список продуктов
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)
	var lastID int = len(products)

	for _, productItem := range products {
		if productItem.ID > lastID {
			lastID = productItem.ID
		}
	}
	newProduct.ID = lastID + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

//Добавление маршрута для получения одного продукта

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновление продукта по id
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем продукт для обновления
	for i, Product := range products {
		if Product.ID == id {

			products[i].ImageURL = updatedProduct.ImageURL
			products[i].Name = updatedProduct.Name
			products[i].Description = updatedProduct.Description
			products[i].Price = updatedProduct.Price

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)           // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)  // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)       // Получить продукт по ID
	http.HandleFunc("/products/update/", updateProductHandler) // Обновить продукт
	http.HandleFunc("/products/delete/", deleteProductHandler) // Удалить продукт

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
