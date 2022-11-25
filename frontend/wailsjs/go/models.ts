export namespace config {
	
	export class Proxy {
	    proxy: string;
	
	    static createFrom(source: any = {}) {
	        return new Proxy(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.proxy = source["proxy"];
	    }
	}
	export class Github {
	    repo: string;
	    branch: string;
	    token: string;
	    path: string;
	    customUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new Github(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.repo = source["repo"];
	        this.branch = source["branch"];
	        this.token = source["token"];
	        this.path = source["path"];
	        this.customUrl = source["customUrl"];
	    }
	}
	export class PicBed {
	    uploader: string;
	    github: Github;
	    proxy: Proxy;
	
	    static createFrom(source: any = {}) {
	        return new PicBed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uploader = source["uploader"];
	        this.github = this.convertValues(source["github"], Github);
	        this.proxy = this.convertValues(source["proxy"], Proxy);
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
	export class Picgo {
	    picBed: PicBed;
	
	    static createFrom(source: any = {}) {
	        return new Picgo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.picBed = this.convertValues(source["picBed"], PicBed);
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
	export class Config {
	    localImagePath: string;
	    enable: boolean;
	    picgo: Picgo;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.localImagePath = source["localImagePath"];
	        this.enable = source["enable"];
	        this.picgo = this.convertValues(source["picgo"], Picgo);
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

export namespace tools {
	
	export class SaveFileArg {
	    totalPath: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new SaveFileArg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.totalPath = source["totalPath"];
	        this.content = source["content"];
	    }
	}
	export class UploadFileArg {
	    content: number[];
	
	    static createFrom(source: any = {}) {
	        return new UploadFileArg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	    }
	}

}

