import type { Pathname } from '$app/types';

export type ShowcaseSection = {
  href: Pathname;
  title: string;
  short: string;
  summary: string;
  description: string;
};

export const showcaseSections: ShowcaseSection[] = [
  {
    href: '/showcase/forms',
    title: 'Forms & controls',
    short: 'fields',
    summary: 'Validation, choices, sliders, switches',
    description:
      'Inputs, labels, validation copy, grouped controls, selection, progress, and submit states.'
  },
  {
    href: '/showcase/layout',
    title: 'Layout & navigation',
    short: 'movement',
    summary: 'Tabs, sidebars, breadcrumbs, pagination',
    description:
      'Navigation, disclosure, tabs, sidebars, pagination, scroll areas, and resizable panes.'
  },
  {
    href: '/showcase/overlays',
    title: 'Overlays & menus',
    short: 'temporary ui',
    summary: 'Dialogs, sheets, command menus, tooltips',
    description:
      'Dialogs, sheets, drawers, menus, popovers, tooltips, command palettes, and toasts.'
  },
  {
    href: '/showcase/display',
    title: 'Display & data',
    short: 'reading',
    summary: 'Tables, calendars, charts, empty states',
    description:
      'Cards, alerts, media, avatars, charts, tables, calendars, empty states, and loading shapes.'
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
