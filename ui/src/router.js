import Vue from 'vue';
import Router from 'vue-router';

import Login from '@/components/Login.vue';
import Auth from '@/components/Auth.vue';
import StreamerView from '@/components/StreamerView.vue';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'login',
            component: Login,
        },
        {
            path: '/redirect',
            name: 'redirect',
            component: Auth,
        },
        {
            path: '/view/:streamerName',
            name: 'view',
            component: StreamerView,
        },
    ],
});
