import {Component, inject} from '@angular/core';
import {Good, GoodsService} from '../header/header.service';
import {App} from '../../app';

@Component({
  selector: 'app-goods',
  imports: [],
  templateUrl: './goods.html',
  styleUrl: './goods.scss',
})
export class Goods {
  readonly goodsService = inject(GoodsService);
  goods: Good[] = [];

  constructor(private app: App) {
    this.goodsService.getAllGoods()
      .subscribe(res => {
        this.goods = res;
      })
  }
}
