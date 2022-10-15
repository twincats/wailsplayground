export namespace models {
	
	export class Manga {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex?: number;
	
	    static createFrom(source: any = {}) {
	        return new Manga(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status_end = source["status_end"];
	        this.mdex = source["mdex"];
	    }
	}

}

export namespace apps {
	
	export class Manga {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Manga(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status_end = source["status_end"];
	        this.mdex = source["mdex"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}

}

