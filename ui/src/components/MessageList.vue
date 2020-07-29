<template>
    <div>
        <iframe
                v-bind:src="'https://player.twitch.tv/?channel=' +streamerName + '&parent=localhost:3000'"
                v-bind:height="400"
                v-bind:width="780"
                v-bind:frameborder="0"
                v-bind:scrolling="no"
                v-bind:allowfullscreen="true">
        </iframe>

        <app-message v-for="event of events"
                     :key="event.id"
                     :message="event">
        </app-message>
    </div>
</template>

<script>
    import gql from 'graphql-tag';
    import Message from '@/components/Message';

    export default {
        components: {
            'app-message': Message,
        },
        data() {
            return {
                events: [],
            };
        },
        props: {
            streamerName: {
                type: String
            }
        },
        created() {
            this.$apollo.subscribe({
                query: gql`
subscription events($name:String!) {
  events (streamerName:$name) {
    id
    subscriberName
    subscribedTo
    type
  }
}
          `,
                variables: {
                    name: this.$props.streamerName
                },
            }).subscribe(
                response => {
                    if (!response.data) {
                        return;
                    }

                    this.$data.events = [response.data.events, ...this.$data.events]
                }
            );
        },
    };
</script>