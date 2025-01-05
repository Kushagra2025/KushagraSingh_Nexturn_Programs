import book_management as bm
import customer_management as cm
import sales_management as sm


while (1):
    print("\nWelcome to BookMarket!\n")
    print("1. Book Management\n2. Customer Management\n3. Sales Management\n4. Exit")
    num = int(input("\nEnter your choice: "))
    if (num == 1):
        print("1. Add Book\n2. View All Books\n3. Search a book")
        user_input = int(input("Enter your Choice: "))
        if (user_input == 1):
            title = input("Enter title: ")
            author = input("Enter author: ")
            price = input("Enter price: ")
            quantity = int(input("Enter quantity: "))
            bm.add_book(title, author, price, quantity)
            
        elif (user_input == 2):
            bm.view_Book()
        
        elif (user_input == 3):
            ask = input('Enter the book name you want to search: ')
            if (bm.search_book(ask)):
                for booki in bm.books:
                    if (booki.title == ask):
                        print("\nBook Details:")
                        print(booki.book_details())
        else :
            continue
    
    elif (num == 2):
        print("1. Add Customer\n2. View All Customers")
        user_input = int(input("Enter your choice: "))
        if (user_input == 1):
            name = input("Enter the Name: ")
            email = input("enter the email: ")
            phone = input("Enter the phone: ")
            cm.add_customer(name, email, phone)
        
        elif (user_input == 2):
            cm.view_customers()
        
        else:
            continue
    
    elif (num == 3):
        print("1. Sell Book\n2. View Sales Record")
        user_input = int(input("Enter your choice: "))
        if (user_input == 1):
            name = input("Enter Customer name: ")
            title = input("Enter title of the book: ")
            sold = int(input("Enter the quantity to sell: "))
            sm.sell_book(name, title, sold)
        
        elif (user_input == 2):
            sm.View_Sales_Record()
        
        else:
            continue
    
    else:
        break