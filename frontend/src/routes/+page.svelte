<script lang="ts">
    import { onMount } from "svelte";

    type Question = {
        question: string;
        correct_answer: string;
        incorrect_answers: string[];
    };

    let questions = $state<Question[] | null>(null);
    let isError = $state(false);

    onMount(() => fetchQuestions());

    function fetchQuestions() {
        fetch("https://opentdb.com/api.php?amount=10&encode=url3986")
            .then((res) => res.json())
            .then((res) => {
                if (res.response_code != 0) {
                    isError = true;
                    return;
                }

                questions = res.results;
            })
            .catch((e) => {
                console.error(e);
                isError = true;
            });
    }
</script>

<svelte:head>
    <title>TurboTrivia</title>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" />
    <link
        href="https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300..800;1,300..800&family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&family=Roboto:ital,wght@0,100..900;1,100..900&display=swap"
        rel="stylesheet"
    />
</svelte:head>

<h1>TurboTrivia</h1>
{#if isError}
    <p>An error occurred. Please try again.</p>
{:else if questions === null}
    <p>Loading...</p>
{:else}
    {#each questions as question}
        <p>{decodeURIComponent(question.question)}</p>
    {/each}
{/if}

<style>
    :global(html, body) {
        color: white;
        font-family: 'Roboto';
    }

    :global(body) {
        background-color: rgb(23, 23, 23);
    }
</style>
