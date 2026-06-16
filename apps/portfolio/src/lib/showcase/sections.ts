import type { Pathname } from '$app/types';

export type ShowcaseSection = {
	href: Pathname;
	title: string;
	label: string;
	description: string;
	count: string;
};

export const showcaseSections: ShowcaseSection[] = [
	{
		href: '/showcase/forms',
		title: 'Forms & controls',
		label: 'inputs',
		description:
			'Form fields, control states, grouped inputs, selection, progress, and command-like buttons.',
		count: '22 families'
	},
	{
		href: '/showcase/layout',
		title: 'Layout & navigation',
		label: 'structure',
		description:
			'Navigation, disclosure, tabs, sidebars, pagination, scroll containers, and resizable panels.',
		count: '10 families'
	},
	{
		href: '/showcase/overlays',
		title: 'Overlays & menus',
		label: 'portals',
		description:
			'Dialogs, sheets, drawers, menus, popovers, tooltips, command palettes, and toasts.',
		count: '12 families'
	},
	{
		href: '/showcase/display',
		title: 'Display & data',
		label: 'surfaces',
		description:
			'Cards, alerts, media, avatars, charts, tables, calendars, empty states, and loading states.',
		count: '12 families'
	}
];

export const representedFamilies = [
	'accordion',
	'alert',
	'alert-dialog',
	'aspect-ratio',
	'avatar',
	'badge',
	'breadcrumb',
	'button',
	'button-group',
	'calendar',
	'card',
	'carousel',
	'chart',
	'checkbox',
	'collapsible',
	'command',
	'context-menu',
	'dialog',
	'drawer',
	'dropdown-menu',
	'empty',
	'field',
	'form',
	'hover-card',
	'input',
	'input-group',
	'input-otp',
	'item',
	'kbd',
	'label',
	'menubar',
	'native-select',
	'navigation-menu',
	'pagination',
	'popover',
	'progress',
	'radio-group',
	'range-calendar',
	'resizable',
	'scroll-area',
	'select',
	'separator',
	'sheet',
	'sidebar',
	'skeleton',
	'slider',
	'sonner',
	'spinner',
	'switch',
	'table',
	'tabs',
	'textarea',
	'toggle',
	'toggle-group',
	'tooltip'
];
