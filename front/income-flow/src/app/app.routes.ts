import { Routes } from '@angular/router';
import { Goods } from './components/goods/goods';
import { App } from './app';
import { BusinessOperations } from './components/business-operations/business-operations';
import { Income } from './components/income/income';
import { Outflow } from './components/outflow/outflow';
import { Remains } from './components/remains/remains';

export const routes: Routes = [
  { path: '', component: App },
  { path: 'goods', component: Goods },
  { path: 'business-operations', component: BusinessOperations },
  { path: 'income', component: Income },
  { path: 'outflow', component: Outflow },
  { path: 'remains', component: Remains },
];
