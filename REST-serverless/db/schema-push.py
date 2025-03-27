import os
from dotenv import load_dotenv
import psycopg2

# Load the environment variables from a parent folder
load_dotenv("../.env")  # Adjust the path if necessary

# Get the DATABASE_URL
DATABASE_URL = os.getenv("DATABASE_URL")

if not DATABASE_URL:
    print("Error: DATABASE_URL is not set.")
    exit(1)

try:
    conn = psycopg2.connect(DATABASE_URL)
    cur = conn.cursor()

    # Read and execute the SQL file
    with open("schema.sql", "r") as f:
        cur.execute(f.read())

    conn.commit()
    cur.close()
    conn.close()
    print("Schema executed successfully.")

except Exception as e:
    print("Error:", e)

