import Vue from 'vue';
import {ApolloClient} from 'apollo-client';
import {HttpLink} from 'apollo-link-http';
import {InMemoryCache} from 'apollo-cache-inmemory';
import VueApollo from 'vue-apollo';
import {split} from 'apollo-link';
import {WebSocketLink} from 'apollo-link-ws';
import {getMainDefinition} from 'apollo-utilities';
import 'bootstrap/scss/bootstrap.scss';
import VueFriendlyIframe from 'vue-friendly-iframe';

import router from './router';
import App from './App.vue';

Vue.config.productionTip = false;

const schema = "http";
const host = "localhost:8080/query";

const httpLink = new HttpLink({
    uri: schema + "://" + host,
    headers: {
        Authorization: localStorage.getItem('token'),
    }
});

const wsLink = new WebSocketLink({
    uri: "ws://" + host,
    options: {
        reconnect: true,
        connectionParams: {
            Authorization: localStorage.getItem('token'),
        }
    }
});

const link = split(
    ({query}) => {
        const {kind, operation} = getMainDefinition(query);
        return kind === 'OperationDefinition' && operation === 'subscription';
    },
    wsLink,
    httpLink,
);
const apolloClient = new ApolloClient({
    link: link,
    cache: new InMemoryCache(),
});
const apolloProvider = new VueApollo({
    defaultClient: apolloClient,
});

Vue.use(VueApollo);
Vue.use(VueFriendlyIframe);

const vm = new Vue({
    router,
    provide: apolloProvider.provide(),
    render: (h) => h(App),
});
vm.$mount('#app');
