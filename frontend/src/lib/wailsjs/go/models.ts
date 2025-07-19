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

}

