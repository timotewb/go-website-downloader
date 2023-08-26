export namespace lib {
	
	export class ResponseType {
	    code: number;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ResponseType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	    }
	}

}

