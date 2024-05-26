export namespace types {
	
	export class Category {
	    id: number;
	    description: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class CategoryInput {
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new CategoryInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	    }
	}
	export class Earning {
	    id: number;
	    description: string;
	    amount: number;
	    profile_id: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Earning(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.profile_id = source["profile_id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class EarningInput {
	    description: string;
	    amount: number;
	    profile_id: number;
	
	    static createFrom(source: any = {}) {
	        return new EarningInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.profile_id = source["profile_id"];
	    }
	}
	export class EarningOutput {
	    id: number;
	    description: string;
	    amount: number;
	    sub_total: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new EarningOutput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.sub_total = source["sub_total"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class EarningToUpdate {
	    description: string;
	    amount: number;
	
	    static createFrom(source: any = {}) {
	        return new EarningToUpdate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.amount = source["amount"];
	    }
	}
	export class EarningUpdate {
	    description: string;
	    amount: number;
	
	    static createFrom(source: any = {}) {
	        return new EarningUpdate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.amount = source["amount"];
	    }
	}
	export class Expense {
	    id: number;
	    description: string;
	    amount: number;
	    profile_id: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Expense(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.profile_id = source["profile_id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class ExpenseInput {
	    description: string;
	    amount: number;
	    category_id: number;
	    profile_id: number;
	
	    static createFrom(source: any = {}) {
	        return new ExpenseInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.category_id = source["category_id"];
	        this.profile_id = source["profile_id"];
	    }
	}
	export class ExpenseOutput {
	    id: number;
	    description: string;
	    amount: number;
	    category: string;
	    sub_total: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new ExpenseOutput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.category = source["category"];
	        this.sub_total = source["sub_total"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class ExpenseToUpdate {
	    description: string;
	    amount: number;
	    category: string;
	
	    static createFrom(source: any = {}) {
	        return new ExpenseToUpdate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.category = source["category"];
	    }
	}
	export class ExpenseUpdate {
	    description: string;
	    amount: number;
	    category_id: number;
	
	    static createFrom(source: any = {}) {
	        return new ExpenseUpdate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.category_id = source["category_id"];
	    }
	}
	export class PaginationOutput {
	    page: number;
	    size: number;
	    total_pages: number;
	    total_items: number;
	    next_page: number;
	    prev_page: number;
	    order_by: string;
	    sort_by: string;
	
	    static createFrom(source: any = {}) {
	        return new PaginationOutput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.total_pages = source["total_pages"];
	        this.total_items = source["total_items"];
	        this.next_page = source["next_page"];
	        this.prev_page = source["prev_page"];
	        this.order_by = source["order_by"];
	        this.sort_by = source["sort_by"];
	    }
	}
	export class PaginatedResult {
	    pagination: PaginationOutput;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new PaginatedResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pagination = this.convertValues(source["pagination"], PaginationOutput);
	        this.data = source["data"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Pagination {
	    page: number;
	    size: number;
	    order_by: string;
	    sort_by: string;
	
	    static createFrom(source: any = {}) {
	        return new Pagination(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.order_by = source["order_by"];
	        this.sort_by = source["sort_by"];
	    }
	}
	
	export class Profile {
	    id: number;
	    description: string;
	    is_default: boolean;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.is_default = source["is_default"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class ProfileInput {
	    description: string;
	    is_default: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ProfileInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.is_default = source["is_default"];
	    }
	}

}

