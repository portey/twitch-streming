<template>
    <div class="row">
        <div class="col-10">
            <div class="row">
                <a :href="link">Login</a>
            </div>
        </div>
    </div>
</template>

<script>
    import gql from 'graphql-tag';

    export default {
        data() {
            return {
                link: "",
            };
        },
        apollo: {
            link() {
                return {
                    query: gql`
          query {
            link: getSignedUrl
          }
        `,
                };
            },
        },
    };
</script>