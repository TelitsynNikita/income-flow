import {Component, signal, inject} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import { form, Field } from '@angular/forms/signals'
import {IncomeService} from './income.service';

@Component({
  selector: 'app-income',
  imports: [MatButtonModule, Field],
  templateUrl: './income.html',
  styleUrl: './income.scss',
})
export class Income {
  incomeForm = signal({
    goods_id: 0,
    section_id: 0,
    goods_count: 0,
    contractors_id: 0
  })
  form = form(this.incomeForm)
  readonly incomeService = inject(IncomeService);

  income() {
    const form = this.incomeForm();

    this.incomeService.income(form.goods_id, form.goods_count, form.contractors_id, form.section_id)
      .subscribe(income => {
        console.log(income)
      })
  }
}
