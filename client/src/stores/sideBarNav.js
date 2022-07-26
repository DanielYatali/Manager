import { writable } from 'svelte/store';

const currentNav = writable('dashboard');

export default currentNav;
