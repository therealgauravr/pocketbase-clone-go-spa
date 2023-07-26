
export const ssr = false;

/** @type {import('./$types').PageLoad} */
export async function load(event) {
    const result = await event.fetch("/hello.json");
    const json = await result.json();

    return { status: 200, ...json };
}

