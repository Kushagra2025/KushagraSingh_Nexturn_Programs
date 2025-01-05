from flask import Flask, request, jsonify
import sqlite3
from datetime import datetime

app = Flask(__name__)

# Helper function to connect to the database
def get_db_connection():
    conn = sqlite3.connect('bookbuddy.db')
    conn.row_factory = sqlite3.Row  # To fetch rows as dictionaries
    return conn

# Helper function to validate genre
valid_genres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]

def validate_book_data(data):
    # Validate genre
    if 'genre' not in data or data['genre'] not in valid_genres:
        return "Invalid genre. Valid genres are: Fiction, Non-Fiction, Mystery, Sci-Fi."
    
    # Validate published year (should be a valid year)
    if 'published_year' not in data or not isinstance(data['published_year'], int):
        return "Invalid year. Published year must be an integer."
    
    # Validate required fields
    required_fields = ['title', 'author', 'published_year', 'genre']
    for field in required_fields:
        if field not in data:
            return f"Missing required field: {field}"
    
    return None  # No errors

# Route to add a new book (POST /books)
@app.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    
    # Validate input data
    validation_error = validate_book_data(data)
    if validation_error:
        return jsonify({"error": "Invalid data", "message": validation_error}), 400
    
    # Insert the book into the database
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('''
    INSERT INTO books (title, author, published_year, genre)
    VALUES (?, ?, ?, ?)
    ''', (data['title'], data['author'], data['published_year'], data['genre']))
    conn.commit()
    
    book_id = cursor.lastrowid
    conn.close()
    
    return jsonify({"message": "Book added successfully", "book_id": book_id}), 201

# Route to get all books (GET /books)
@app.route('/books', methods=['GET'])
def get_all_books():
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM books')
    books = cursor.fetchall()
    conn.close()
    
    # Convert to list of dictionaries
    books_list = [dict(book) for book in books]
    
    return jsonify(books_list), 200

# Route to get a specific book by ID (GET /books/<id>)
@app.route('/books/<int:id>', methods=['GET'])
def get_book(id):
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM books WHERE id = ?', (id,))
    book = cursor.fetchone()
    conn.close()
    
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    
    return jsonify(dict(book)), 200

# Route to update a book (PUT /books/<id>)
@app.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    data = request.get_json()
    
    # Validate input data
    validation_error = validate_book_data(data)
    if validation_error:
        return jsonify({"error": "Invalid data", "message": validation_error}), 400
    
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('''
    UPDATE books SET title = ?, author = ?, published_year = ?, genre = ?
    WHERE id = ?
    ''', (data['title'], data['author'], data['published_year'], data['genre'], id))
    
    conn.commit()
    
    if cursor.rowcount == 0:
        conn.close()
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    
    conn.close()
    return jsonify({"message": "Book updated successfully"}), 200

# Route to delete a book (DELETE /books/<id>)
@app.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('DELETE FROM books WHERE id = ?', (id,))
    conn.commit()
    
    if cursor.rowcount == 0:
        conn.close()
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    
    conn.close()
    return jsonify({"message": "Book deleted successfully"}), 200

# Run the Flask application
if __name__ == '__main__':
    app.run(debug=True)
