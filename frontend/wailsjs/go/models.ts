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
	
	export class CheckJobIsRunningType {
	    job_count: number;
	
	    static createFrom(source: any = {}) {
	        return new CheckJobIsRunningType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.job_count = source["job_count"];
	    }
	}

}

