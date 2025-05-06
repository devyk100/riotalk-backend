import os
import psycopg2
import sys
from dotenv import load_dotenv

# Load environment variables
load_dotenv("../.env")  # Adjust the path if necessary
DATABASE_URL = os.getenv("DATABASE_URL")

if not DATABASE_URL:
    print("Error: DATABASE_URL is not set.")
    sys.exit(1)

def connect_db():
    """Establish a database connection."""
    try:
        return psycopg2.connect(DATABASE_URL)
    except Exception as e:
        print("Database connection error:", e)
        sys.exit(1)

def drop_database():
    """Drop all tables, sequences, types, and enums in the database."""
    conn = connect_db()
    cur = conn.cursor()
    try:
        # Drop all tables
        cur.execute("""
            DO $$ DECLARE
                r RECORD;
            BEGIN
                FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
                    EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
                END LOOP;
            END $$;
        """)

        # Drop all sequences
        cur.execute("""
            DO $$ DECLARE
                r RECORD;
            BEGIN
                FOR r IN (SELECT sequence_name FROM information_schema.sequences WHERE sequence_schema = 'public') LOOP
                    EXECUTE 'DROP SEQUENCE IF EXISTS ' || quote_ident(r.sequence_name) || ' CASCADE';
                END LOOP;
            END $$;
        """)

        # Drop all types (includes enums)
        cur.execute("""
DO $$ DECLARE
                r RECORD;
            BEGIN
                FOR r IN (SELECT typname FROM pg_type
                          WHERE typnamespace = (SELECT oid FROM pg_namespace WHERE nspname = 'public')
                          AND typname NOT LIKE '\\_%') LOOP  -- Skip array types
                    EXECUTE 'DROP TYPE IF EXISTS ' || quote_ident(r.typname) || ' CASCADE';
                END LOOP;
            END $$;

        """)

        conn.commit()
        print("✅ Database reset: All tables, sequences, and types dropped successfully.")
    except Exception as e:
        print("Error dropping database objects:", e)
    finally:
        cur.close()
        conn.close()

def run_queries():
    """Execute SQL queries from the schema.sql file."""
    conn = connect_db()
    cur = conn.cursor()
    try:
        with open("schema.sql", "r") as f:
            sql = f.read()
            cur.execute(sql)
        conn.commit()
        print("✅ Queries executed successfully.")
    except Exception as e:
        print("Error executing queries:", e)
    finally:
        cur.close()
        conn.close()

def show_data():
    """Fetch and display all data from all tables in a readable format."""
    conn = connect_db()
    cur = conn.cursor()
    try:
        cur.execute("SELECT tablename FROM pg_tables WHERE schemaname = 'public';")
        tables = cur.fetchall()
        if not tables:
            print("⚠️ No tables found in the database.")
            return

        for table in tables:
            table_name = table[0]
            print(f"\n📌 Data from table: {table_name}")
            print("-" * 50)
            cur.execute(f"SELECT * FROM {table_name};")
            rows = cur.fetchall()
            if not rows:
                print("⚠️ No data in this table.")
            else:
                for row in rows:
                    print(row)
            print("-" * 50)

    except Exception as e:
        print("Error retrieving data:", e)
    finally:
        cur.close()
        conn.close()

# Interactive Menu
def main():
    while True:
        print("\n📌 Database Management Menu")
        print("1️⃣ Drop the entire database (⚠️ Dangerous! Includes tables, sequences, and types)")
        print("2️⃣ Run queries from schema.sql")
        print("3️⃣ Show all data in the database")
        print("4️⃣ Exit")

        choice = input("Enter your choice (1-4): ").strip()

        if choice == "1":
            confirm = input("⚠️ Are you sure you want to drop everything? (yes/no): ").strip().lower()
            if confirm == "yes":
                drop_database()
            else:
                print("❌ Operation canceled.")
        elif choice == "2":
            run_queries()
        elif choice == "3":
            show_data()
        elif choice == "4":
            print("👋 Exiting program.")
            sys.exit(0)
        else:
            print("❌ Invalid choice. Please enter a number between 1-4.")

if __name__ == "__main__":
    main()
