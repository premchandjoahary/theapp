# Theapp

This project is regarding loan management system. 

## Features

- Users can apply for loans by providing the "amount required" and "loan term."
- Loan applications assume a "weekly" repayment frequency.
- Admins can approve loan applications, transitioning them from "PENDING" to "APPROVED."
- Authenticated users can submit weekly loan repayments, marking them as "PAID."
- The system automatically updates the loan status to "PAID" when all scheduled repayments are "PAID."

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

Make sure you have the following software installed:

- [Go](https://golang.org/dl/) (1.11 or higher)
- [Git](https://git-scm.com/)

### Installing

1. Clone the repository:

```bash
git clone https://github.com/premchandjoahary/theapp.git
cd theapp
run make
