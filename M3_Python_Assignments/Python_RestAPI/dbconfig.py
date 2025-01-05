import sqlite3

def init_db():
    # Connect to SQLite database (or create it if it doesn't exist)
    conn = sqlite3.connect('bookbuddy.db')
    cursor = conn.cursor()
    
    # Create the 'books' table
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        published_year INTEGER NOT NULL,
        genre TEXT NOT NULL
    );
    ''')
    
    # Insert sample data
    cursor.executemany('''
    INSERT INTO books (title, author, published_year, genre) 
    VALUES (?, ?, ?, ?);
    ''', [
        ('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'Fiction'),
        ('To Kill a Mockingbird', 'Harper Lee', 1960, 'Fiction'),
        ('1984', 'George Orwell', 1949, 'Sci-Fi')
    ])
    
    # Commit and close the connection
    conn.commit()
    conn.close()

# Run the function to initialize the database
if __name__ == "__main__":
    init_db()
    print("Database and sample data initialized.")
