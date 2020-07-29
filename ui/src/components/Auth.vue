<template>
    <div class="row">
        <div class="col-10">
            <div class="row">
                <form v-on:submit.prevent="onSubmit">
                    <div class="measure">
                        <label htmlFor="name">Streamer name and press enter</label>
                        <input id="name" type="text" required v-model="streamerName"/>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import gql from 'graphql-tag';
    import Router from '../router'

    export default {
        data() {
            return {
                streamerName: "",
                token: "",
            };
        },
        watch: {
            token: {
                handler() {
                    localStorage.setItem('token', "Bearer " + this.token);
                },
                deep: true,
            },
        },
        methods: {
            onSubmit() {
                let localStreamerName = this.streamerName
                this.$apollo.mutate({
                    mutation: gql`
mutation subscribe($name:String!) {
  subscribeToStreamer(name:$name)
}
`,
                    variables: {
                        name: this.streamerName
                    },
                    awaitRefetchQueries: true,
                    update: () => {
                        Router.push({name: 'view', params: {streamerName: localStreamerName}})
                    }
                })

                this.streamerName = "";
            }
        },
        apollo: {
            token() {
                return {
                    query: gql`
query exchange($token:String!) {
  token:exchangeToken(accessToken:$token)
}
        `,
                    variables: {
                        token: this.$route.query.access_token
                    }

                };
            },
        },
    };
</script>