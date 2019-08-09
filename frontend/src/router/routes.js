import Layout from '../pages/Layout';
import Home from '../pages/Home';
import About from '../pages/About';

export default [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '',
                name: 'home',
                component: Home
            },
            {
                path: 'about',
                name: 'about',
                component: About
            }
        ]
    }
]
