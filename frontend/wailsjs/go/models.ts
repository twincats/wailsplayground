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
	export class MangaHomeApi {
	    id: number;
	    title: string;
	    status_end: boolean;
	    mdex?: number;
	    chapter_id: number;
	    chapter: string;
	    // Go type: sqltime.Time
	    download_time: any;
	
	    static createFrom(source: any = {}) {
	        return new MangaHomeApi(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status_end = source["status_end"];
	        this.mdex = source["mdex"];
	        this.chapter_id = source["chapter_id"];
	        this.chapter = source["chapter"];
	        this.download_time = this.convertValues(source["download_time"], null);
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
	export class MangaHome {
	    manga: MangaHomeApi[];
	    // Go type: Pagination
	    pagination?: any;
	
	    static createFrom(source: any = {}) {
	        return new MangaHome(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manga = this.convertValues(source["manga"], MangaHomeApi);
	        this.pagination = this.convertValues(source["pagination"], null);
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

