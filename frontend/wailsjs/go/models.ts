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
	    session_id: string;
	    stale_flag: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ActivityDataType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.favicon_url = source["favicon_url"];
	        this.session_id = source["session_id"];
	        this.stale_flag = source["stale_flag"];
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
	export class GalleryType {
	    site_name: string;
	    site_location: string;
	    favicon: string;
	
	    static createFrom(source: any = {}) {
	        return new GalleryType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.site_name = source["site_name"];
	        this.site_location = source["site_location"];
	        this.favicon = source["favicon"];
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
	export class SettingsType {
	    content_dir: string;
	    content_dir_exists: boolean;
	    content_dir_wsport: number;
	
	    static createFrom(source: any = {}) {
	        return new SettingsType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content_dir = source["content_dir"];
	        this.content_dir_exists = source["content_dir_exists"];
	        this.content_dir_wsport = source["content_dir_wsport"];
	    }
	}

}

