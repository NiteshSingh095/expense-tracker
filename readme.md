# 💰 Expense Tracker CLI (Go)

A simple yet powerful command-line expense tracker built using Go. This project helps users manage their daily expenses efficiently with persistent storage and clean CLI interactions.

---

## 🚀 Features

* ➕ Add expenses with description, amount, and category
* 📋 List all expenses in a clean tabular format
* ❌ Delete expenses by ID
* ✏️ Update existing expenses
* 📊 View total and monthly expense summaries
* ⚠️ Budget warning system
* 📁 Export expenses to CSV

---

## 🛠️ Tech Stack

* Language: Go
* Data Storage: JSON file
* CLI Handling: flag package

---

Project Url : https://roadmap.sh/projects/expense-tracker

## 📦 Project Structure

expense-tracker/
│── main.go
│── cmd/
│    ├── add.go
│    ├── list.go
│    ├── delete.go
│    ├── update.go
│    ├── summary.go
│    ├── export.go
│
│── internal/
│    ├── expense.go
│    ├── storage.go
│
│── data/
│    └── expenses.json

---

## ▶️ Usage

### Add Expense

```
expense-tracker add --description "Lunch" --amount 200 --category food
```

### List Expenses

```
expense-tracker list
```

### Delete Expense

```
expense-tracker delete --id 1
```

### Update Expense

```
expense-tracker update --id 1 --amount 300
```

### Summary

```
expense-tracker summary
expense-tracker summary --month 3
```

### Export CSV

```
expense-tracker export
```

---

## 💡 Learnings

* CLI application design in Go
* Structs and JSON data handling
* File I/O operations
* Modular code architecture
* Real-world problem solving

---

## 📌 Future Improvements

* REST API version
* GUI using Flutter
* Cloud sync
* Authentication system

---

## ⭐ If you like this project, give it a star!
