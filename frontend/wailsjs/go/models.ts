export namespace main {
	
	export class CityAutocomplete {
	    city_uuid: number[];
	    code: number;
	    full_name: string;
	    country_code: string;
	
	    static createFrom(source: any = {}) {
	        return new CityAutocomplete(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.city_uuid = source["city_uuid"];
	        this.code = source["code"];
	        this.full_name = source["full_name"];
	        this.country_code = source["country_code"];
	    }
	}
	export class Location {
	    country_code: string;
	    region_code: number;
	    region: string;
	    city_code: number;
	    city: string;
	    fias_guid: number[];
	    postal_code: string;
	    longitude: number;
	    latitude: number;
	    address: string;
	    address_full: string;
	    city_uuid: number[];
	
	    static createFrom(source: any = {}) {
	        return new Location(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.country_code = source["country_code"];
	        this.region_code = source["region_code"];
	        this.region = source["region"];
	        this.city_code = source["city_code"];
	        this.city = source["city"];
	        this.fias_guid = source["fias_guid"];
	        this.postal_code = source["postal_code"];
	        this.longitude = source["longitude"];
	        this.latitude = source["latitude"];
	        this.address = source["address"];
	        this.address_full = source["address_full"];
	        this.city_uuid = source["city_uuid"];
	    }
	}
	export class Warning {
	    code: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Warning(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	    }
	}
	export class Error {
	    code: string;
	    additional_code: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Error(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.additional_code = source["additional_code"];
	        this.message = source["message"];
	    }
	}
	export class Dimension {
	    width: number;
	    height: number;
	    depth: number;
	
	    static createFrom(source: any = {}) {
	        return new Dimension(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.width = source["width"];
	        this.height = source["height"];
	        this.depth = source["depth"];
	    }
	}
	export class Timeend {
	    hour: number;
	    minute: number;
	    second: number;
	    nano: number;
	
	    static createFrom(source: any = {}) {
	        return new Timeend(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hour = source["hour"];
	        this.minute = source["minute"];
	        this.second = source["second"];
	        this.nano = source["nano"];
	    }
	}
	export class Timestart {
	    hour: number;
	    minute: number;
	    second: number;
	    nano: number;
	
	    static createFrom(source: any = {}) {
	        return new Timestart(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hour = source["hour"];
	        this.minute = source["minute"];
	        this.second = source["second"];
	        this.nano = source["nano"];
	    }
	}
	export class Worktimeexception {
	    date_start: string;
	    date_end: string;
	    time_start: Timestart;
	    time_end: Timeend;
	    is_working: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Worktimeexception(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date_start = source["date_start"];
	        this.date_end = source["date_end"];
	        this.time_start = this.convertValues(source["time_start"], Timestart);
	        this.time_end = this.convertValues(source["time_end"], Timeend);
	        this.is_working = source["is_working"];
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
	export class Worktime {
	    day: number;
	    time: string;
	
	    static createFrom(source: any = {}) {
	        return new Worktime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.day = source["day"];
	        this.time = source["time"];
	    }
	}
	export class Officeimage {
	    number: number;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new Officeimage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.number = source["number"];
	        this.url = source["url"];
	    }
	}
	export class Phone {
	    number: string;
	    additional: string;
	
	    static createFrom(source: any = {}) {
	        return new Phone(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.number = source["number"];
	        this.additional = source["additional"];
	    }
	}
	export class Deliverypoints {
	    code: string;
	    address: string;
	    name: string;
	    uuid: string;
	    address_comment: string;
	    nearest_station: string;
	    nearest_metro_station: string;
	    work_time: string;
	    phones: Phone[];
	    email: string;
	    note: string;
	    type: string;
	    owner_code: string;
	    take_only: boolean;
	    is_handout: boolean;
	    is_reception: boolean;
	    is_dressing_room: boolean;
	    is_marketplace: boolean;
	    is_ltl: boolean;
	    have_cashless: boolean;
	    have_cash: boolean;
	    have_fast_payment_system: boolean;
	    allowed_cod: boolean;
	    site: string;
	    office_image_list: Officeimage[];
	    work_time_list: Worktime[];
	    work_time_exception_list: Worktimeexception[];
	    weight_min: number;
	    weight_max: number;
	    dimensions: Dimension[];
	    errors: Error[];
	    warnings: Warning[];
	    location: Location;
	    distance: number;
	    fulfillment: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Deliverypoints(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.address = source["address"];
	        this.name = source["name"];
	        this.uuid = source["uuid"];
	        this.address_comment = source["address_comment"];
	        this.nearest_station = source["nearest_station"];
	        this.nearest_metro_station = source["nearest_metro_station"];
	        this.work_time = source["work_time"];
	        this.phones = this.convertValues(source["phones"], Phone);
	        this.email = source["email"];
	        this.note = source["note"];
	        this.type = source["type"];
	        this.owner_code = source["owner_code"];
	        this.take_only = source["take_only"];
	        this.is_handout = source["is_handout"];
	        this.is_reception = source["is_reception"];
	        this.is_dressing_room = source["is_dressing_room"];
	        this.is_marketplace = source["is_marketplace"];
	        this.is_ltl = source["is_ltl"];
	        this.have_cashless = source["have_cashless"];
	        this.have_cash = source["have_cash"];
	        this.have_fast_payment_system = source["have_fast_payment_system"];
	        this.allowed_cod = source["allowed_cod"];
	        this.site = source["site"];
	        this.office_image_list = this.convertValues(source["office_image_list"], Officeimage);
	        this.work_time_list = this.convertValues(source["work_time_list"], Worktime);
	        this.work_time_exception_list = this.convertValues(source["work_time_exception_list"], Worktimeexception);
	        this.weight_min = source["weight_min"];
	        this.weight_max = source["weight_max"];
	        this.dimensions = this.convertValues(source["dimensions"], Dimension);
	        this.errors = this.convertValues(source["errors"], Error);
	        this.warnings = this.convertValues(source["warnings"], Warning);
	        this.location = this.convertValues(source["location"], Location);
	        this.distance = source["distance"];
	        this.fulfillment = source["fulfillment"];
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

