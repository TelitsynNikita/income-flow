import {Component, inject} from '@angular/core';
import {BusinessOperation, BusinessOperation2, BusinessOperationsService} from './business-operations.service';
import {App} from '../../app';

@Component({
  selector: 'app-business-operations',
  imports: [],
  templateUrl: './business-operations.html',
  styleUrl: './business-operations.scss',
})
export class BusinessOperations {
  readonly businessOperationForm = inject(BusinessOperationsService);
  operations: BusinessOperation[] = [];
  operations2: {date: string, operation_id: number, values: BusinessOperation2[]}[] = [];

  constructor(private app: App) {
    this.businessOperationForm.getOperations()
      .subscribe(res => {
        this.operations = res;
      })

    this.businessOperationForm.getOperations2().subscribe(res => {
      res.forEach((value, key) => {
        this.operations2?.push({
          date: key,
          operation_id: value[0]?.operation_id,
          values: value
        })
      })
    })
  }
}
