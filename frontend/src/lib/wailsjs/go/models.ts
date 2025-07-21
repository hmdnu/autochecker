export namespace handler {
	
	export class GameResponse {
	    data: any;
	    name: string;
	    message: string;
	    retcode: number;
	
	    static createFrom(source: any = {}) {
	        return new GameResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.name = source["name"];
	        this.message = source["message"];
	        this.retcode = source["retcode"];
	    }
	}

}

export namespace service {
	
	export class Account {
	    Id: string;
	    Name: string;
	    Token: string;
	    Uid: string;
	
	    static createFrom(source: any = {}) {
	        return new Account(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Name = source["Name"];
	        this.Token = source["Token"];
	        this.Uid = source["Uid"];
	    }
	}
	export class Response {
	    Id: string;
	    AccountName: string;
	    DateCheck: string;
	    Game: handler.GameResponse[];
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.AccountName = source["AccountName"];
	        this.DateCheck = source["DateCheck"];
	        this.Game = this.convertValues(source["Game"], handler.GameResponse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

