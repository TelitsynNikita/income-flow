import { Component, inject } from '@angular/core';
import { Remain, RemainsService } from './remains.service';

@Component({
  selector: 'app-remains',
  imports: [],
  templateUrl: './remains.html',
  styleUrl: './remains.scss',
})
export class Remains {
  readonly goodsService = inject(RemainsService);
  remains: Remain[] = [];

  constructor() {
    this.goodsService.getRemains()
      .subscribe(res => {
        this.remains = res;
      })
  }
}
