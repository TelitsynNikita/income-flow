import {Component, signal, inject} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import { form, Field } from '@angular/forms/signals'
import {OutflowService} from './outflow.service';

@Component({
  selector: 'app-outflow',
    imports: [
      MatButtonModule, Field
    ],
  templateUrl: './outflow.html',
  styleUrl: './outflow.scss',
})
export class Outflow {
  outflowForm = signal({
    goods_id: 0,
    section_id: 0,
    goods_count: 0,
    contractors_id: 0
  })
  form = form(this.outflowForm)
  readonly incomeService = inject(OutflowService);

  outflow() {
    const form = this.outflowForm();

    this.incomeService.outflow(form.goods_id, form.goods_count, form.contractors_id, form.section_id)
      .subscribe(outflow => {
        console.log(outflow)
      })
  }
}
