export namespace lib {
	
	export class ResponseType {
	    code: number;
	    message: string;
	    url: string;
	    favicon_url: string;
	
	    static createFrom(source: any = {}) {
	        return new ResponseType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.url = source["url"];
	        this.favicon_url = source["favicon_url"];
	    }
	}

}

export namespace models {
	
	export class ActivityDataType {
	    url: string;
	    favicon_url: string;
	
	    static createFrom(source: any = {}) {
	        return new ActivityDataType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.favicon_url = source["favicon_url"];
	    }
	}
	export class CheckActivityType {
	    job_count: number;
	    data: ActivityDataType[];
	
	    static createFrom(source: any = {}) {
	        return new CheckActivityType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.job_count = source["job_count"];
	        this.data = this.convertValues(source["data"], ActivityDataType);
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
	export class ResponseType {
	    code: number;
	    message: string;
	    url: string;
	    favicon_url: string;
	
	    static createFrom(source: any = {}) {
	        return new ResponseType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.url = source["url"];
	        this.favicon_url = source["favicon_url"];
	    }
	}

}

