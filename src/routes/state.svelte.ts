type Props = {
	[key: string]: unknown;
};

type SidebarItem = {
	title: string;
	active: boolean;
	href?: string;
	props?: Props;
};

export type SidebarData = {
	items: SidebarItem[];
};

export class ApplicationState {
	sidebar: SidebarState;

	constructor(sidebar: SidebarState) {
		this.sidebar = sidebar;
	}
}

export class SidebarState {
	items = $state<SidebarItem[]>([]);

	constructor(data: SidebarData) {
		this.items = data.items;
	}

	setActive(pathname: string) {
		this.#reset();

		this.items.forEach((item) => {
			if (item.href === pathname) {
				item.active = true;
			}
		});
	}

	#reset() {
		this.items.forEach((item) => {
			item.active = false;
		});
	}
}
