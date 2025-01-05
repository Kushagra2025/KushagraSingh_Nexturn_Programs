class Book:
    quantity = 0
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity
        
    def book_details(self):
        return f'Title: {self.title}\nAuthor: {self.author}\nPrice: {self.price}\nQuantity: {self.quantity}'


# book_1 = Book('The Hero', 'Rhonda Byrne', 'Rs. 590', 10)
# print(book_1.book_details())

books = []

def add_book(title, author, price, quantity):
    new_book = Book(title, author, price, quantity)
    books.append(new_book)
    print("\nBook Added Successfully!")


def view_Book():
    for i,booki in enumerate(books):
        print(f'Book {i+1}')
        print(booki.book_details())
        print("----------------------------")


def search_book(name):
    for booki in books:
        if (booki.title == name):
            return True
        else :
            return False

# user_input = input('Enter Y or N:')

# while (user_input == "Y" or user_input == "y"):
#     title = input("Enter title: ")
#     author = input("Enter author: ")
#     price = input("Enter price: ")
#     quantity = int(input("Enter quantity: "))
#     add_book(title, author, price, quantity)

#     user_input = input('Enter Y or N:')

# view_Book()
# ask = input('Enter the book name you want to search: ')
# print(search_book(ask))
