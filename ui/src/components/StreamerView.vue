<template>
    <div class="row">
        <div class="col-10">
            <div class="row">
                <h2>Viewing streamer "{{ streamerName }}"</h2>
                <app-message-list class="col-12" v-bind:streamerName="streamerName"/>
            </div>
        </div>
    </div>
</template>

<script>
    import MessageList from '@/components/MessageList.vue';

    export default {
        data() {
            return {
                streamerName: this.$route.params.streamerName,
            }
        },
        components: {
            'app-message-list': MessageList,
        }
    };
</script>
