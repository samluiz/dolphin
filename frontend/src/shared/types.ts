export interface Earning {
    id: number
    description: string
    amount: number
}

export interface Expense extends Earning {
    category: string
}

export interface EarningInput {
    description: string
    amount: number
}

export interface ExpenseInput extends EarningInput {
    category: string
}

export interface Item {
    description: string
    amount: number
    category?: string
}