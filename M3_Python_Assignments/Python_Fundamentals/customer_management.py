class customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone
        
    def cust_details(self):
         return f'Customer Details:\nname: {self.name}\nemail: {self.email}\nPhone: {self.phone}'
     
# c_1 = customer('John','john@gmail.com', '122-1213')
# print(c_1.cust_details())

customers = []

def add_customer(name, email, phone):
    new_customer = customer(name, email, phone)
    customers.append(new_customer)

    
def view_customers():
    for i, custi in enumerate(customers):
        print(f'Customer {i+1}')
        print(custi.cust_details())
        print('-----------------------------')
    