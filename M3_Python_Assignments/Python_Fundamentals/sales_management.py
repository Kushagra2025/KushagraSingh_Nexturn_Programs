import customer_management as cm
import book_management as bm

class transaction(cm.customer):
    def __init__(self, name, title, sold):
        self.name = name
        self.title = title
        self.sold = sold
        
    def trans_details(self):
        return f'Name: {self.name}\nTitle: {self.title}\nSold: {self.sold}'
    
sales = []

def sell_book(name, title, sold):
    new_trans = transaction(name, title, sold)
    sales.append(new_trans)
    for book in bm.books:
        if (book.title == title):
            book.quantity = book.quantity - sold
            print("Sale Successfull")
            print(f'Quantity left: {book.quantity}\n')
    
def View_Sales_Record():
    for i,transi in enumerate(sales):
        print(f'Sales Record {i+1}')
        print(transi.trans_details())
        print('-----------------------------')

# bm.add_book('Inner Engineering', 'Sadhguru', 'Rs. 490', 20)
# sell_book('Alexa','Inner Engineering', 5)